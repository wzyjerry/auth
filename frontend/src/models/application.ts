import type { ImmerReducer } from 'umi';
import { message } from 'antd';
import type { EffectsCommandMap, Effect } from 'dva';
import {
  CreateRequest,
  CreateReply,
  RetrieveRequest,
  RetrieveReply,
  GenerateClientSecretRequest,
  GenerateClientSecretReply,
  RevokeClientSecretRequest,
  Secret,
  UploadLogoRequest,
  UpdateRequest,
  DeleteRequest,
  GetAllReply,
} from '@/api/application/v1/application';
import { Application } from '@/api/application/v1/application';
import { ApplicationClient } from '@/api/request';
import type { Empty } from '@/api/google/protobuf/empty';
import { history } from 'umi';
export { Application };
export interface ApplicationModelType {
  state: Application & GetAllReply;
  effects: {
    create: (action: { payload: CreateRequest }, effects: EffectsCommandMap) => void;
    update: (action: { payload: Omit<UpdateRequest, 'id'> }, effects: EffectsCommandMap) => void;
    delete: Effect;
    setup: (action: { payload: RetrieveRequest }, effects: EffectsCommandMap) => void;
    setupAll: Effect;
    generateClientSecret: (
      action: { payload: Omit<GenerateClientSecretRequest, 'id'> },
      effects: EffectsCommandMap,
    ) => void;
    revokeClientSecret: (
      action: { payload: Omit<RevokeClientSecretRequest, 'id'> },
      effects: EffectsCommandMap,
    ) => void;
    uploadLogo: (
      action: { payload: Omit<UploadLogoRequest, 'id'> },
      effects: EffectsCommandMap,
    ) => void;
  };
  reducers: {
    set: ImmerReducer<
      Application & GetAllReply,
      {
        type: 'set';
        payload: Application & GetAllReply & { clear?: boolean };
      }
    >;
    pushFrontClientSecret: ImmerReducer<
      Application,
      {
        type: 'pushFrontClientSecret';
        payload: Secret;
      }
    >;
    removeClientSecret: ImmerReducer<
      Application,
      {
        type: 'removeClientSecret';
        payload: string;
      }
    >;
  };
  subscriptions: unknown;
}

// 如果首个Secret not masked, mask it.
const adjustClientSecrets = (secrets: Secret[]): Secret[] => {
  if (secrets.length > 0) {
    if (!secrets[0].masked) {
      secrets[0].masked = true;
      secrets[0].secret = `*****${secrets[0].secret.slice(secrets[0].secret.length - 5)}`;
    }
  }
  return secrets;
};

const ApplicationModel: ApplicationModelType = {
  state: Object.assign(Application.fromJSON({}), GetAllReply.fromJSON({})),

  effects: {
    *create({ payload }, { call }) {
      const { response, error }: { response: CreateReply; error: Error } = yield call(
        ApplicationClient.Create,
        payload,
      );
      if (response) {
        history.push(`/settings/applications/${response.id}`);
      } else if (error) {
        throw error;
      }
    },
    *update({ payload }, { call, put, select }) {
      const { id } = yield select(({ application }: { application: Application }) => application);
      const request: UpdateRequest = { id, ...payload };
      const { response, error }: { response: GenerateClientSecretReply; error: Error } = yield call(
        ApplicationClient.Update,
        request,
      );
      if (response) {
        yield put({
          type: 'set',
          payload: {
            ...payload,
            clear: false,
          },
        });
        message.success('应用更新成功');
      } else if (error) {
        throw error;
      }
    },
    *delete(_, { call, select }) {
      const { id } = yield select(({ application }: { application: Application }) => application);
      const request: DeleteRequest = { id };
      const { response, error }: { response: Empty; error: Error } = yield call(
        ApplicationClient.Delete,
        request,
      );
      if (response) {
        history.push(`/settings/applications`);
      } else if (error) {
        throw error;
      }
    },
    *setup({ payload }, { call, put }) {
      const { response, error }: { response: RetrieveReply; error: Error } = yield call(
        ApplicationClient.Retrieve,
        payload,
      );
      const application: Application = {
        id: payload.id,
        ...response,
      };
      if (response) {
        yield put({
          type: 'set',
          payload: {
            ...application,
            clear: true,
          },
        });
      } else if (error) {
        throw error;
      }
    },
    *setupAll(_, { call, put }) {
      const { response, error }: { response: GetAllReply; error: Error } = yield call(
        ApplicationClient.GetAll,
      );
      if (response) {
        yield put({
          type: 'set',
          payload: {
            ...response,
            clear: true,
          },
        });
      } else if (error) {
        throw error;
      }
    },
    *generateClientSecret({ payload }, { call, put, select }) {
      const { id } = yield select(({ application }: { application: Application }) => application);
      const request: GenerateClientSecretRequest = { id, ...payload };
      const { response, error }: { response: GenerateClientSecretReply; error: Error } = yield call(
        ApplicationClient.GenerateClientSecret,
        request,
      );
      if (response) {
        yield put({
          type: 'pushFrontClientSecret',
          payload: response.secret,
        });
      } else if (error) {
        throw error;
      }
    },
    *revokeClientSecret({ payload }, { call, put, select }) {
      const { id } = yield select(({ application }: { application: Application }) => application);
      const request: RevokeClientSecretRequest = { id, ...payload };
      const { response, error }: { response: Empty; error: Error } = yield call(
        ApplicationClient.RevokeClientSecret,
        request,
      );
      if (response) {
        yield put({
          type: 'removeClientSecret',
          payload: payload.secretId,
        });
      } else if (error) {
        throw error;
      }
    },
    *uploadLogo({ payload }, { call, put, select }) {
      const { id } = yield select(({ application }: { application: Application }) => application);
      const request: UploadLogoRequest = { id, ...payload };
      const { response, error }: { response: Empty; error: Error } = yield call(
        ApplicationClient.UploadLogo,
        request,
      );
      if (response) {
        yield put({
          type: 'set',
          payload: {
            logo: payload.logo,
            clear: false,
          },
        });
      } else if (error) {
        throw error;
      }
    },
  },
  reducers: {
    set(state, { payload }) {
      const { clear } = payload;
      delete payload.clear;
      if (clear) {
        Object.assign(state, Object.assign(Application.fromJSON({}), GetAllReply.fromJSON({})));
      }
      Object.assign(state, payload);
    },
    pushFrontClientSecret(state, { payload }) {
      state.clientSecrets = adjustClientSecrets(state.clientSecrets);
      state.clientSecrets = [payload, ...state.clientSecrets];
    },
    removeClientSecret(state, { payload }) {
      state.clientSecrets = adjustClientSecrets(state.clientSecrets).filter(
        (secret) => secret.id != payload,
      );
    },
  },
  subscriptions: {},
};

export default ApplicationModel;
