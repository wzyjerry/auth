import style from './index.less';
import { LoadingOutlined, PlusOutlined } from '@ant-design/icons';
import { Upload, message } from 'antd';
import type { UploadChangeParam } from 'antd/es/upload';
import type { RcFile, UploadFile, UploadProps } from 'antd/es/upload/interface';
import { useState } from 'react';
import { useDispatch, useSelector } from 'umi';
import type { Application } from 'umi';
import type { UploadLogoRequest } from '@/api/application/v1/application';

const getBase64 = (img: RcFile, callback: (url: string) => void) => {
  const reader = new FileReader();
  reader.addEventListener('load', () => callback(reader.result as string));
  reader.readAsDataURL(img);
};

const beforeUpload = (file: RcFile) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
  if (!isJpgOrPng) {
    message.error('您只能上传JPG/PNG格式的文件!');
  }
  const isLt2M = file.size / 1024 / 1024 < 2;
  if (!isLt2M) {
    message.error('Logo大小必须小于2MB!');
  }
  return isJpgOrPng && isLt2M;
};

const Logo: React.FC = () => {
  const avatar = useSelector(({ application }: { application: Application }) => application.logo);
  const dispatch = useDispatch();
  const [loading, setLoading] = useState(false);
  const handleChange: UploadProps['onChange'] = (info: UploadChangeParam<UploadFile>) => {
    if (info.file.status === 'uploading') {
      setLoading(true);
      return;
    }
    if (info.file.status === 'done') {
      getBase64(info.file.originFileObj as RcFile, (url) => {
        setLoading(false);
        const request: Omit<UploadLogoRequest, 'id'> = {
          logo: url,
        };
        dispatch({
          type: 'application/uploadLogo',
          payload: request,
        });
      });
    }
  };
  const uploadButton = (
    <div>
      {loading ? <LoadingOutlined /> : <PlusOutlined />}
      <div className={style.uploadLabel}>上传Logo</div>
    </div>
  );
  return (
    <Upload
      listType="picture-card"
      className={style.logo}
      showUploadList={false}
      beforeUpload={beforeUpload}
      onChange={handleChange}
    >
      {avatar ? <img src={avatar} alt="avatar" className={style.image} /> : uploadButton}
    </Upload>
  );
};
export default Logo;
