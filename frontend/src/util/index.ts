import type { ResponseError } from 'umi-request';
import { message } from 'antd';
import { history } from 'umi';
export class Err extends Error {
  constructor(name: string, message?: string) {
    super(message);
    this.name = name;
  }
}
interface backendError {
  code: number;
  reason: string;
  message: string;
  metadata: Map<string, string>;
}

export function errorHandler(error: ResponseError<backendError>) {
  if (error.data) {
    switch (error.data.reason) {
      case 'UNAUTHORIZED':
      case 'FORBIDDEN':
        history.push(
          `/user/login?return_to=${encodeURIComponent(
            location.pathname + location.search + location.hash,
          )}`,
        );
        return;
      case 'APPLICATION_NOT_FOUND':
        history.push(`/404`);
        return;
      default:
        return Promise.reject(new Err(error.data.reason, error.data.message));
    }
  }
  message.error(`[NETWORK_ERROR] ${error}`);
}
