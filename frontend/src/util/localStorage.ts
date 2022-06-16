import { Resolve, Reject } from '.';
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

export const ErrSetLocalstorage = new Err('SetLocalstorageError');
export const ErrGetLocalstorage = new Err('GetLocalstorageError');
export const ErrKeyNotExists = new Err('KeyNotExistsError');

class helper {
  Save(key: string, value: string): void | Reject {
    try {
      localStorage.setItem(key, value);
    } catch {
      return new Reject(ErrSetLocalstorage);
    }
  }
  Load(key: string): Resolve<string> | Reject {
    try {
      const value = localStorage.getItem(key);
      if (value === null) {
        return new Reject(ErrKeyNotExists);
      }
      return new Resolve(value);
    } catch {
      return new Reject(ErrGetLocalstorage);
    }
  }
}

export class Token extends helper {
  static AUTH_TOKEN_KEY = 'AUTH_TOKEN';
  Save(token: string): void | Reject {
    return super.Save(Token.AUTH_TOKEN_KEY, token)
  }
  Load(): Resolve<string> | Reject<Error> {
      return super.Load(Token.AUTH_TOKEN_KEY)
  }
}

export class Avatar extends helper {
  static AVATAR_KEY = 'AVATAR';
  Save(avatar: string): void | Reject {
    return super.Save(Avatar.AVATAR_KEY, avatar)
  }
  Load(): Resolve<string> | Reject<Error> {
      return super.Load(Avatar.AVATAR_KEY)
  }
}
