import type { ResponseError } from 'umi-request';
import { message } from 'antd';
import { history } from 'umi';
import { Err } from './localStorage';

export function Go<T, E = Error>(promise: Promise<T>): Promise<Resolve<T> | Reject<E>> {
  return promise
    .then<Resolve<T>>((data: T) => new Resolve(data))
    .catch((err: E) => new Reject(err));
}

export class Resolve<T> {
  val: T;
  constructor(val: T) {
    this.val = val;
  }
}

export class Reject<T = Error> {
  error: T;
  constructor(error: T) {
    this.error = error;
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
