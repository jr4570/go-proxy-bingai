export interface ApiResult<T> {
  code: ApiResultCode;
  message: string;
  data: T;
}

export enum ApiResultCode {
  /**
   * 成功
   */
  OK = 200,

  /**
   * 未授權
   */
  Unauthorized = 401,
}
