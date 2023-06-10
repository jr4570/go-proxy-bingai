<script setup lang="ts">
import { ref } from 'vue';
import { NModal, NList, NListItem, NButton, useMessage, NSpace, NInput, NUpload, type UploadFileInfo, NEmpty } from 'naive-ui';
import { usePromptStore, type IPrompt, type IPromptDownloadConfig } from '@/stores/modules/prompt';
import { storeToRefs } from 'pinia';
import VirtualList from 'vue3-virtual-scroll-list';
import ChatPromptItem from './ChatPromptItem.vue';

const messgae = useMessage();
const promptStore = usePromptStore();
const { promptDownloadConfig, isShowPromptSotre, promptList, keyword, searchPromptList, optPromptConfig } = storeToRefs(promptStore);

const isShowDownloadPop = ref(false);

const isImporting = ref(false);
const isExporting = ref(false);

const showAddPromptPop = () => {
  optPromptConfig.value.isShow = true;
  optPromptConfig.value.type = 'add';
  optPromptConfig.value.title = '添加提示詞';
  optPromptConfig.value.newPrompt = {
    act: '',
    prompt: '',
  };
};

const savePrompt = () => {
  const { type, tmpPrompt, newPrompt } = optPromptConfig.value;
  if (!newPrompt.act) {
    return messgae.error('提示詞標題不能為空');
  }
  if (!newPrompt.prompt) {
    return messgae.error('提示詞描述不能為空');
  }
  if (type === 'add') {
    promptList.value = [newPrompt, ...promptList.value];
    messgae.success('添加提示詞成功');
  } else if (type === 'edit') {
    if (newPrompt.act === tmpPrompt?.act && newPrompt.prompt === tmpPrompt?.prompt) {
      messgae.warning('提示詞未變更');
      optPromptConfig.value.isShow = false;
      return;
    }
    const rawIndex = promptList.value.findIndex((x) => x.act === tmpPrompt?.act && x.prompt === tmpPrompt?.prompt);
    if (rawIndex > -1) {
      promptList.value[rawIndex] = newPrompt;
      messgae.success('編輯提示詞成功');
    } else {
      messgae.error('編輯提示詞出錯');
    }
  }
  optPromptConfig.value.isShow = false;
};

const readFile = (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = function (ev) {
      resolve(ev.target?.result as string);
    };
    reader.onerror = reject;
    reader.readAsText(file);
  });
};

const importPrompt = async (options: { file: UploadFileInfo; fileList: Array<UploadFileInfo>; event?: Event }) => {
  // console.log(options.file);
  if (options.file.file) {
    isImporting.value = true;
    const fileText = await readFile(options.file.file);
    const promptData = JSON.parse(fileText);
    const result = promptStore.addPrompt(promptData);
    if (result.result) {
      messgae.info(`上傳文件含 ${promptData.length} 條數據`);
      messgae.success(`成功導入 ${result.data?.successCount} 條有效數據`);
    } else {
      messgae.error(result.msg || '提示詞格式有誤');
    }
    isImporting.value = false;
  } else {
    messgae.error('上傳文件有誤');
  }
};

const exportPrompt = () => {
  if (promptList.value.length === 0) {
    return messgae.error('暫無可導出的提示詞數據');
  }
  isExporting.value = true;
  const jsonDataStr = JSON.stringify(promptList.value);
  const blob = new Blob([jsonDataStr], { type: 'application/json' });
  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.download = 'BingAIPrompts.json';
  link.click();
  URL.revokeObjectURL(url);
  messgae.success('導出提示詞庫成功');
  isExporting.value = false;
};

const clearPrompt = () => {
  promptList.value = [];
  messgae.success('清空提示詞庫成功');
};

const downloadPrompt = async (config: IPromptDownloadConfig) => {
  if (!config.url) {
    return messgae.error('請先輸入下載鏈接');
  }
  config.isDownloading = true;
  let jsonData: Array<IPrompt>;
  if (config.url.endsWith('.json')) {
    jsonData = await fetch(config.url).then((res) => res.json());
  } else if (config.url.endsWith('.csv')) {
    const csvData = await fetch(config.url).then((res) => res.text());
    console.log(csvData);
    jsonData = csvData
      .split('\n')
      .filter((x) => x)
      .map((x) => {
        const arr = x.split('","');
        return {
          act: arr[0].slice(1),
          prompt: arr[1]?.slice(1),
        };
      });
    jsonData.shift();
  } else {
    config.isDownloading = false;
    return messgae.error('暫不支持下載此後綴的提示詞');
  }
  config.isDownloading = false;
  const result = promptStore.addPrompt(jsonData);
  if (result.result) {
    messgae.info(`下載文件含 ${jsonData.length} 條數據`);
    messgae.success(`成功導入 ${result.data?.successCount} 條有效數據`);
  } else {
    messgae.error(result.msg || '提示詞格式有誤');
  }
};
</script>

<template>
  <NModal class="w-11/12 xl:w-[900px]" v-model:show="isShowPromptSotre" preset="card" title="提示詞庫">
    <div class="flex justify-start flex-wrap gap-2 px-5 pb-2">
      <NInput class="basis-full xl:basis-0 xl:min-w-[300px]" placeholder="搜索提示詞" v-model:value="keyword" :clearable="true"></NInput>
      <NButton secondary type="info" @click="isShowDownloadPop = true">下載</NButton>
      <NButton secondary type="info" @click="showAddPromptPop">添加</NButton>
      <NUpload class="w-[56px] xl:w-auto" accept=".json" :default-upload="false" :show-file-list="false" @change="importPrompt">
        <NButton secondary type="success" :loading="isImporting">導入</NButton>
      </NUpload>
      <!-- <NButton secondary type="success">導入</NButton> -->
      <NButton secondary type="success" @click="exportPrompt" :loading="isExporting">導出</NButton>
      <NButton secondary type="error" @click="clearPrompt">清空</NButton>
    </div>
    <VirtualList
      v-if="searchPromptList.length > 0"
      class="h-[40vh] xl:h-[60vh] overflow-y-auto"
      :data-key="'prompt'"
      :data-sources="searchPromptList"
      :data-component="ChatPromptItem"
      :keeps="10"
    />
    <NEmpty v-else class="h-[40vh] xl:h-[60vh] flex justify-center items-center" description="暫無數據">
      <template #extra>
        <NButton secondary type="info" @click="isShowDownloadPop = true">下載提示詞</NButton>
      </template>
    </NEmpty>
  </NModal>
  <NModal class="w-11/12 xl:w-[600px]" v-model:show="optPromptConfig.isShow" preset="card" :title="optPromptConfig.title">
    <NSpace vertical>
      標題
      <NInput placeholder="請輸入標題" v-model:value="optPromptConfig.newPrompt.act"></NInput>
      描述
      <NInput placeholder="請輸入描述" type="textarea" v-model:value="optPromptConfig.newPrompt.prompt"></NInput>
      <NButton block secondary type="info" @click="savePrompt">保存</NButton>
    </NSpace>
  </NModal>
  <NModal class="w-11/12 xl:w-[600px]" v-model:show="isShowDownloadPop" preset="card" title="下載提示詞">
    <NList class="overflow-y-auto rounded-lg" hoverable clickable>
      <NListItem v-for="(config, index) in promptDownloadConfig" :key="index">
        <a v-if="config.type === 1" class="no-underline text-blue-500" :href="config.url" target="_blank" rel="noopener noreferrer">{{ config.name }}</a>
        <NInput v-else-if="config.type === 2" placeholder="請輸入下載鏈接，支持 json 及 csv " v-model:value="config.url"></NInput>
        <template #suffix>
          <div class="flex justify-center gap-5">
            <a class="no-underline" v-if="config.type === 1" :href="config.refer" target="_blank" rel="noopener noreferrer">
              <NButton secondary>來源</NButton>
            </a>
            <NButton secondary type="info" @click="downloadPrompt(config)" :loading="config.isDownloading">下載</NButton>
          </div>
        </template>
      </NListItem>
    </NList>
  </NModal>
</template>
