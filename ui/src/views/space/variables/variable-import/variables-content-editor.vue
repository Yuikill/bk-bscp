<template>
  <Teleport :disabled="!isOpenFullScreen" to="body">
    <div :class="['config-content-editor', { fullscreen: isOpenFullScreen }]">
      <div class="editor-title">
        <div class="tips">
          <InfoLine class="info-icon" />
          {{ t('仅支持大小不超过') }} 5M
        </div>
        <div class="btns">
          <i
            class="bk-bscp-icon icon-separator"
            v-bk-tooltips="{
              content: t('分隔符'),
              placement: 'top',
              distance: 20,
            }"
            @click="separatorShow = !separatorShow" />
          <Search
            v-bk-tooltips="{
              content: t('搜索'),
              placement: 'top',
              distance: 20,
            }"
            @click="codeEditorRef.openSearch()" />
          <bk-upload
            :accept="format === 'text' ? '.txt' : `.${format}`"
            theme="button"
            :size="5"
            :custom-request="handleUploadFile">
            <template #trigger>
              <Upload
                v-bk-tooltips="{
                  content: t('上传'),
                  placement: 'top',
                  distance: 20,
                }" />
            </template>
          </bk-upload>
          <i
            :class="['bk-bscp-icon', 'icon-terminal', { isOpen: modelValue }]"
            v-bk-tooltips="{
              content: t('示例面板'),
              placement: 'top',
              distance: 20,
            }"
            @click="emits('update:modelValue', !modelValue)" />
          <FilliscreenLine
            v-if="!isOpenFullScreen"
            v-bk-tooltips="{
              content: t('全屏'),
              placement: 'top',
              distance: 20,
            }"
            @click="handleOpenFullScreen" />
          <UnfullScreen
            v-else
            v-bk-tooltips="{
              content: t('退出全屏'),
              placement: 'bottom',
              distance: 20,
            }"
            @click="handleCloseFullScreen" />
        </div>
      </div>
      <div class="editor-content">
        <CodeEditor
          ref="codeEditorRef"
          :model-value="editorContent"
          :error-line="errorLine"
          :language="format"
          @paste="handlePaste"
          @enter="separatorShow = true"
          @validate="handleValidateEditor"
          @update:model-value="handleContentChange" />
        <div class="separator" v-show="separatorShow">
          <SeparatorSelect @closed="separatorShow = false" @confirm="handleSelectSeparator" />
        </div>
        <slot name="sufContent" :fullscreen="isOpenFullScreen"></slot>
      </div>
    </div>
  </Teleport>
</template>
<script setup lang="ts">
  import { ref, onBeforeUnmount, watch, computed, nextTick } from 'vue';
  import { useI18n } from 'vue-i18n';
  import BkMessage from 'bkui-vue/lib/message';
  import { InfoLine, FilliscreenLine, UnfullScreen, Search, Upload } from 'bkui-vue/lib/icon';
  import { importVariablesText, importVariablesJSON, importVariablesYaml } from '../../../../api/variable';
  import CodeEditor from '../../../../components/code-editor/index.vue';
  import SeparatorSelect from './separator-select.vue';
  import useGlobalStore from '../../../../store/global';
  import { storeToRefs } from 'pinia';

  interface errorLineItem {
    lineNumber: number;
    errorInfo: string;
  }

  const { t } = useI18n();

  const { spaceId } = storeToRefs(useGlobalStore());

  const props = defineProps<{
    modelValue: boolean;
    format: string;
  }>();

  const emits = defineEmits(['hasError', 'update:modelValue']);

  const isOpenFullScreen = ref(false);
  const codeEditorRef = ref();
  const separatorShow = ref(false);
  const textContent = ref('');
  const separator = ref(' ');
  const shouldValidate = ref(false);
  const errorLine = ref<errorLineItem[]>([]);
  const jsonContent = ref('');
  const yamlContent = ref('');

  const editorContent = computed(() => {
    if (props.format === 'text') return textContent.value;
    if (props.format === 'json') return jsonContent.value;
    return yamlContent.value;
  });

  watch(
    () => editorContent.value,
    (val) => {
      if (!val) {
        emits('hasError', true);
        return;
      }
      if (props.format === 'text') {
        handleValidateEditor();
      } else {
        nextTick(() => emits('hasError', !codeEditorRef.value.validate(val)));
      }
    },
    { immediate: true },
  );

  watch(
    () => errorLine.value,
    (val) => {
      shouldValidate.value = val.length > 0;
    },
  );

  onBeforeUnmount(() => {
    codeEditorRef.value.destroy();
  });
  // 打开全屏
  const handleOpenFullScreen = () => {
    isOpenFullScreen.value = true;
    window.addEventListener('keydown', handleEscClose, { once: true });
    BkMessage({
      theme: 'primary',
      message: t('按 Esc 即可退出全屏模式'),
    });
  };

  const handleCloseFullScreen = () => {
    isOpenFullScreen.value = false;
    window.removeEventListener('keydown', handleEscClose);
  };

  // Esc按键事件处理
  const handleEscClose = (event: KeyboardEvent) => {
    if (event.code === 'Escape') {
      isOpenFullScreen.value = false;
    }
  };

  const handleContentChange = (val: string) => {
    if (props.format === 'text') {
      textContent.value = val;
    } else if (props.format === 'json') {
      jsonContent.value = val;
    } else {
      yamlContent.value = val;
    }
  };

  // 校验编辑器内容
  const handleValidateEditor = () => {
    const textContentArray = textContent.value.split('\n').map((item) => item.trim());
    errorLine.value = [];
    let hasSeparatorError = false;
    textContentArray.forEach((item, index) => {
      if (item === '') return;
      const regex = separator.value === ' ' ? /\s+/ : separator.value;
      const textContentContent = item.split(regex).map((item) => item.trim());
      const key = textContentContent[0];
      const type = textContentContent[1];
      const value = textContentContent[2];
      value === '""' && (textContentContent[2] = ''); // "" 转空字符串 代表变量为空值
      if (textContentContent.length < 3) {
        errorLine.value.push({
          errorInfo: t('请检查是否已正确使用分隔符'),
          lineNumber: index + 1,
        });
        hasSeparatorError = true;
      } else if (!key.startsWith('bk_bscp_') && !key.startsWith('BK_BSCP_')) {
        errorLine.value.push({
          errorInfo: t('变量必须以bk_bscp_或BK_BSCP_开头'),
          lineNumber: index + 1,
        });
      } else if (type !== 'string' && type !== 'number') {
        errorLine.value.push({
          errorInfo: t('类型必须为 string 或者 number'),
          lineNumber: index + 1,
        });
      } else if (type === 'number' && !/^\d+(\.\d+)?$/.test(value)) {
        errorLine.value.push({
          errorInfo: t('类型为number 值不为number'),
          lineNumber: index + 1,
        });
      }
    });
    emits('hasError', textContent.value && errorLine.value.length > 0);
    return hasSeparatorError;
  };

  // 导入kv
  const handleImport = async () => {
    let res;
    if (props.format === 'text') {
      const params = {
        separator: separator.value === ' ' ? 'white-space' : separator.value,
        variables: editorContent.value,
      };
      res = await importVariablesText(spaceId.value, params);
    } else if (props.format === 'json') {
      res = await importVariablesJSON(spaceId.value, editorContent.value);
    } else {
      res = await importVariablesYaml(spaceId.value, editorContent.value);
    }
    return res.data.ids;
  };

  const handleSelectSeparator = (selectSeparator: string) => {
    separator.value = selectSeparator;
    handleValidateEditor();
  };

  const handleUploadFile = async (option: { file: File }) => {
    const reader = new FileReader();
    reader.readAsText(option.file);
    reader.onload = function (e) {
      const fileContent = e.target?.result as string;
      if (props.format === 'text') {
        textContent.value = fileContent;
      } else if (props.format === 'json') {
        jsonContent.value = fileContent;
      } else {
        yamlContent.value = fileContent;
      }
    };
  };

  const handlePaste = () => {
    if (handleValidateEditor()) separatorShow.value = true;
  };

  defineExpose({
    handleImport,
  });
</script>
<style lang="scss" scoped>
  .config-content-editor {
    height: 640px;
    padding-top: 10px;
    &.fullscreen {
      position: fixed;
      top: 0;
      left: 0;
      width: 100vw;
      height: 100vh;
      z-index: 5000;
    }
    .editor-title {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 0 16px;
      height: 40px;
      color: #979ba5;
      background: #2e2e2e;
      border-radius: 2px 2px 0 0;
      .tips {
        display: flex;
        align-items: center;
        font-size: 12px;
        .info-icon {
          margin-right: 4px;
        }
      }
      .btns {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 8px;
        span,
        i {
          display: flex;
          align-items: center;
          justify-content: center;
          width: 24px;
          height: 24px;
          border-radius: 2px;
          font-size: 16px;
          color: #979ba5;
          cursor: pointer;
          &:hover {
            color: #3a84ff;
          }
          &.isOpen {
            background: #181818;
          }
        }
        :deep(.bk-upload) {
          display: flex;
          justify-content: space-between;
          align-items: center;
          .bk-upload-list {
            display: none;
          }
        }
      }
    }
    .editor-content {
      position: relative;
      height: calc(100% - 130px);
      .separator {
        position: absolute;
        right: 0;
        top: 0;
      }
    }
  }
</style>
