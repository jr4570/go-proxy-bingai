import { ref } from 'vue';
import { defineStore } from 'pinia';

export interface SydneyConfig {
  baseUrl: string;
  label: string;
  isUsable?: boolean;
  delay?: number;
  isCus?: boolean;
}

export interface CheckSydneyConfigResult {
  isUsable: boolean;
  errorMsg?: string;
  delay?: number;
}

export const useChatStore = defineStore(
  'chat-store',
  () => {
    const chatHubPath = '/sydney/ChatHub';

    const isShowChatServiceSelectModal = ref(false);
    const selectedSydneyBaseUrl = ref('');
    const sydneyConfigs = ref<SydneyConfig[]>([
      {
        baseUrl: 'https://sydney.bing.com',
        label: 'Bing 官方',
      },
      {
        baseUrl: 'https://sydney.vcanbb.chat',
        label: 'Cloudflare',
      },
      {
        baseUrl: location.origin,
        label: '本站',
      },
      {
        baseUrl: '',
        label: '自定義',
        isCus: true,
      },
    ]);
    const sydneyCheckTimeoutMS = 3000;

    const checkSydneyConfig = async (config: SydneyConfig): Promise<CheckSydneyConfigResult> => {
      if (!config.baseUrl) {
        return {
          isUsable: false,
          errorMsg: '鏈接不可為空',
        };
      }
      try {
        const startTime = Date.now();
        const ws = new WebSocket(config.baseUrl.replace('http', 'ws') + chatHubPath);
        const wsTimer = setTimeout(() => {
          ws.close();
        }, sydneyCheckTimeoutMS);
        await new Promise((resolve, reject) => {
          ws.onopen = () => {
            clearTimeout(wsTimer);
            resolve(ws.close());
          };
          ws.onerror = () => {
            clearTimeout(wsTimer);
            reject(new Error(`聊天服務器 ${config.baseUrl} 連接失敗`));
          };
          ws.onclose = () => reject(new Error(`聊天服務器 ${config.baseUrl} 連接超時`));
        });
        return {
          isUsable: true,
          delay: Date.now() - startTime,
        };
      } catch (error) {
        return {
          isUsable: false,
          errorMsg: error instanceof Error ? error.message : '',
        };
      }
    };

    const checkAllSydneyConfig = async () => {
      const checkAllConfigPromises = sydneyConfigs.value
        .filter((x) => x.baseUrl)
        .map(async (config) => {
          const checkResult = await checkSydneyConfig(config);
          config.isUsable = checkResult.isUsable;
          config.delay = checkResult.delay;
        });
      await Promise.all(checkAllConfigPromises);
    };

    return {
      isShowChatServiceSelectModal,
      sydneyConfigs,
      selectedSydneyBaseUrl,
      checkSydneyConfig,
      checkAllSydneyConfig,
    };
  },
  {
    persist: {
      key: 'chat-store',
      storage: localStorage,
      paths: ['selectedSydneyBaseUrl', 'sydneyConfigs'],
    },
  }
);
