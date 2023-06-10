declare const sj_evt: {
  bind: (n: string, t: Function, i: boolean, r?: any) => void;
  fire: (n: string) => void;
};
declare const SydneyFullScreenConv: {
  initWithWaitlistUpdate: (n: object, t: number) => void;
};

/**
 * 更有創造力 | 更平衡 | 更精確
 */
type ToneType = 'Creative' | 'Balanced' | 'Precise';

interface BingMessage {
  /**
   * 消息內容
   */
  text: string;

  /**
   * 還可以圖片鏈接？
   */
  imageUrl?: string;
}

type BingMessageType = keyof {
  ActionRequest: 'ActionRequest';
  Ads: 'Ads';
  AdsQuery: 'AdsQuery';
  CaptchaChallenge: 'CaptchaChallenge';
  Chat: 'Chat';
  Context: 'Context';
  Disengaged: 'Disengaged';
  SearchQuery: 'SearchQuery';
  Internal: 'Internal';
  Suggestion: 'Suggestion';
  InternalSuggestions: 'InternalSuggestions';
  InternalSearchQuery: 'InternalSearchQuery';
  InternalSearchResult: 'InternalSearchResult';
  InternalActionMarker: 'InternalActionMarker';
  InternalStateMarker: 'InternalStateMarker';
  InternalLoaderMessage: 'InternalLoaderMessage';
  Progress: 'Progress';
  GenerateContentQuery: 'GenerateContentQuery';
  RenderCardRequest: 'RenderCardRequest';
  SemanticSerp: 'SemanticSerp';
  Any: 'Any';
  ChatName: 'ChatName';
};

interface SuggestedResponses {
  contentOrigin: string;
  hiddenText?: string;
  messageId: string;
  messageType: BingMessageType;
  offense: string;
  text: string;
}
interface TextMessageModel {
  author: string;
  messageId: string;
  messageType: BingMessageType;
  type: string;
  contentOrigin: string;
  imageUrl: string;
  isFinalized: boolean;
  text: string;
  suggestedResponses: SuggestedResponses[];
}

interface BingChat {
  // 是否請求響應中
  isRequestPending: boolean;
  api: {
    bing: {
      _endpoint: string;
      captcha: {
        client: {
          sendOperationRequest: (operationArguments: object, operationSpec: object) => {};
        };
      };
      conversation: {
        /**
         * 創建請求
         */
        create: (O) => {};
        /**
         * 聊天記錄
         */
        getChats: (O) => {};
      };
    };
    sydney: {};
  };
  requestToken: {
    cancel: () => Promise<unknown>;
    /**
     * 聊天完成
     * @param O
     * @returns
    */
    complete: (O) => Promise<unknown>;
    error: (O) => Promise<unknown>;
  };
  onConversationExpired: PublicSubscribeEvent;
  onLoadConversationInvoked: PublicSubscribeEvent;
  onMessage: PublicSubscribeEvent;
  onPendingRequestStateChanged: PublicSubscribeEvent;
  onRequestGenerated: PublicSubscribeEvent;
  onResponseRendered: PublicSubscribeEvent;
  onStreamingComplete: PublicSubscribeEvent;
}

interface BingConversation {
  clientId: string;
  conversationType: string;
  hashedSignature: string;
  id: string;
  isExpired: boolean;
  messages: TextMessageModel[];
  state: string;
  suggestions: SuggestedResponses[];
  updateId: (Id: string, expiry: Date, clientId: string, signature: string) => {};
}

type PublicSubscribeEvent = (callback: Function, thisArgs, disposables) => {};

declare const CIB: {
  /**
   * 微軟 bing 版本信息
   */
  version: {
    buildTimestamp: string;
    commit: string;
    version: string;
  };
  /**
   * 整個必應聊天 cib-serp
   */
  vm: {
    errorService: {
      onAppErrorStateChange: PublicSubscribeEvent;
      onChatErrorStateChange: PublicSubscribeEvent;
      onChatWarningStateChange: PublicSubscribeEvent;
    };
    /**
     * 是否手機版
     */
    isMobile: boolean;
    actionBar: {
      /**
       * 輸入框
       */
      input: HTMLTextAreaElement;
      /**
       * 輸入框文本 賦值即輸入問題
       */
      inputText: string;
      /**
       * 自動建議的前置文本
       */
      autoSuggestPrependedText: string;
      /**
       * 自動建議附加文本
       */
      autoSuggestAppendedText: string;
      /**
       * 提交當前輸入框問題
       */
      submitInputText: () => {};
    };
    conversation: BingConversation;
    /**
     * 歷史記錄
     */
    sidePanel: {
      /**
       * M 是否顯示
       */
      isVisibleMobile: boolean;
      /**
       * PC 是否顯示  get shouldShowPanel
       */
      isVisibleDesktop: boolean;
    };
    /**
     * 選擇對話樣式
     */
    toneSelector: {
      /**
       * 更有創造力 | 更平衡 | 更精確  -- 直接賦值即可切換 set tone(O)
       */
      tone: ToneType;
    };
  };
  config: {
    bing: {
      baseUrl: string;
    };
    edgeAction: {
   /**
       * hook 接收消息，需啟用
       */
      hookIncomingMessage: (message) => boolean;
      isEnabled: boolean;
    };
    features: {
      enableThreads: boolean;
      /**
       * 獲取聊天歷史
       */
      enableGetChats: boolean;
    };
    sydney: {
      baseUrl: string;
      /**
       * 安全域名？移除 localhost，開發即可 create
       */
      hostnamesToBypassSecureConnection: string[];
      expiryInMinutes: number;
    };
    messaging: {
      /**
       * 單次最大對話數
       */
      maxTurnsPerConversation: number;
      /**
       * 打字機速度調節，默認 1000 / 15 = 55 （每秒15字？）
       */
      messageBufferWorkerStreamDelayMS: number;
    };
  };
  manager: {
    chat: BingChat;
    conversation: BingConversation;
    /**
     * 重置聊天
     */
    resetConversation: () => {};

    /**
     * 發送消息
     * @param O 消息內容
     * @param B 默認 false ，則發送消息
     * @param G 默認 chat 消息類型
     * @param U 默認 Keyboard 似乎只是區分輸入及語音
     * @returns
     */
    sendMessage: (O: BingMessage, B?: boolean, G?: BingMessageType, U?: 'Keyboard' | 'Speech') => {};

    onResetConversationInvoked: PublicSubscribeEvent;
  };

  onConsentGiven: PublicSubscribeEvent;
  onConversationExpired: PublicSubscribeEvent;
  onConversationRequestStateChange: PublicSubscribeEvent;
  onInputMethodChanged: PublicSubscribeEvent;
  onMobileUpsellPopupShown: PublicSubscribeEvent;
  onModeChanged: PublicSubscribeEvent;
  onModeChanging: PublicSubscribeEvent;
  onResetConversation: PublicSubscribeEvent;
  onResponseRendered: PublicSubscribeEvent;
  onResponseToneChanged: PublicSubscribeEvent;
  onSerpSlotSuggestionInvoked: PublicSubscribeEvent;
  /**
   * 接收流完成
   */
  onStreamingComplete: PublicSubscribeEvent;
  onThreadLoadInvoked(listener);
  onThreadLoaded(listener);
  onWorkToggleChanged: PublicSubscribeEvent;

  responseTone: ToneType;
};
