import { Err, Resolve, Reject } from '.';

export const ErrSetLocalstorage = new Err('SetLocalstorageError');
export const ErrGetLocalstorage = new Err('GetLocalstorageError');
export const ErrTokenNotExists = new Err('TokenNotExistsError');

export class Token {
  static AUTH_TOKEN_KEY = 'AUTH_TOKEN_KEY';
  Save(token: string): void | Reject {
    try {
      localStorage.setItem(Token.AUTH_TOKEN_KEY, token);
    } catch {
      return new Reject(ErrSetLocalstorage);
    }
  }
  Load(): Resolve<string> | Reject {
    try {
      const token = localStorage.getItem(Token.AUTH_TOKEN_KEY);
      if (token === null) {
        return new Reject(ErrTokenNotExists);
      }
      return new Resolve(token);
    } catch {
      return new Reject(ErrGetLocalstorage);
    }
  }
}
