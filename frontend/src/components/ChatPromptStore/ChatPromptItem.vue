<script setup lang="ts">
import { usePromptStore, type IPrompt } from '@/stores/modules/prompt';
import { NThing, NTag, NButton, NEllipsis, useMessage } from 'naive-ui';
import { storeToRefs } from 'pinia';

defineProps<{
  index: number;
  source: IPrompt;
}>();

const messgae = useMessage();
const promptStore = usePromptStore();
const { promptList, optPromptConfig } = storeToRefs(promptStore);

const delPrompt = (item: IPrompt) => {
  promptList.value = promptList.value.filter((x) => x.act !== item.act && x.prompt !== item.prompt);
  messgae.success('刪除提示詞成功');
};

const showEditPromptPop = (item: IPrompt) => {
  optPromptConfig.value.isShow = true;
  optPromptConfig.value.type = 'edit';
  optPromptConfig.value.title = '編輯提示詞';
  optPromptConfig.value.tmpPrompt = item;
  optPromptConfig.value.newPrompt = { ...item };
};
</script>

<template>
  <NThing class="hover:bg-gray-100 cursor-pointer p-5">
    <template #description>
      <NTag type="info">
        <span class="inline-block max-w-[120px] xl:max-w-[650px] overflow-hidden text-ellipsis">{{ source.act }}</span>
      </NTag>
      <div class="float-right">
        <NButton secondary type="info" size="small" @click="showEditPromptPop(source)">編輯</NButton>
        <NButton secondary class="ml-2" type="error" size="small" @click="delPrompt(source)">刪除</NButton>
      </div>
    </template>
    <NEllipsis :tooltip="false" :line-clamp="2">{{ source.prompt }}</NEllipsis>
  </NThing>
</template>
