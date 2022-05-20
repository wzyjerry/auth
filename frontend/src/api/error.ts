import { notification } from 'antd';
import { ResponseError } from 'umi-request';

interface error {
  code: number;
  reason: string;
  message: string;
  metadata: Map<string, string>;
}

const errorHandler = (error: ResponseError<error>) => {
  if (error.response) {
    notification.error({
      message: `请求错误 ${error.data.code}: ${error.data.reason}`,
      description: error.data.message,
    });
  } else {
    notification.error({
      message: `网络错误`,
      description: error.message,
    });
  }
  return Promise.reject(error);
};

export default errorHandler;
