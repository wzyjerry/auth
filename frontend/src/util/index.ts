import { ResponseError } from 'umi-request';

export function Go<T, E = Error>(
  promise: Promise<T>,
): Promise<Resolve<T> | Reject<E>> {
  return promise
    .then<Resolve<T>>((data: T) => new Resolve(data))
    .catch((err: E) => new Reject(err));
}

export class Err extends Error {
  constructor(name: string, message?: string) {
    super(message);
    this.name = name;
  }
}

export class Throw extends Err {
  constructor(error: Err) {
    super(error.name, error.message);
  }
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
    return Promise.reject(new Err(`[${error.data.code}]${error.data.reason}`, error.data.message))
  }
  return Promise.reject(new Err("NETWORD_ERROR", "netword error"))
};
