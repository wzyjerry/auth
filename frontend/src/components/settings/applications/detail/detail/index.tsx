import style from './index.less';
import copy from 'copy-to-clipboard';
import { CopyOutlined } from '@ant-design/icons';
import { Button, Alert, message } from 'antd';
import { useState } from 'react';
import SecretItem from './secret';
import Logo from './logo';
import UpdateForm from './update';
import GenerateClientSecret from './generateClientSecret';
import type { Application } from 'umi';
import { useSelector } from 'umi';

const Detail: React.FC = () => {
  const [visible, setVisible] = useState(false);
  const application: Application = useSelector(
    ({ application }: { application: Application }) => application,
  );
  const copyClientId = (): void => {
    copy(application.clientId);
    message.success(`已复制到剪切板: ${application.clientId}`);
  };
  const renderClientSecrets = () => {
    if (application.clientSecrets.length > 0) {
      const only = application.clientSecrets.length === 1;
      return (
        <div className={style.clientSecretArea}>
          {application.clientSecrets[0].masked || (
            <Alert
              className={style.alert}
              type="info"
              description="确保现在复制您的新客户端密钥。你将无法再次看到它。"
            />
          )}
          <div className={style.clientSecretList}>
            {application.clientSecrets.map((secret) => (
              <SecretItem key={secret.id} secret={secret} only={only} />
            ))}
          </div>
        </div>
      );
    } else {
      return <span>您需要一个客户端密钥来验证应用程序。</span>;
    }
  };
  return (
    <div className={style.detail}>
      <div className={style.title}>
        <span>{application.name}</span>
      </div>
      <hr className={style.line} />
      <div className={style.clientIdLine}>
        <span className={style.clientIdLabel}>客户端ID:</span>
        <span className={style.clientId}>{application.clientId}</span>
        <Button
          type="link"
          shape="circle"
          icon={<CopyOutlined />}
          size="large"
          onClick={copyClientId}
        />
      </div>
      <div className={style.clientSecretLine}>
        <span className={style.clientSecretLabel}>客户端密码</span>
        <Button className={style.generate} size="large" onClick={() => setVisible(true)}>
          生成新的客户端密码
        </Button>
        <GenerateClientSecret
          id={application.id}
          visiable={visible}
          onClose={() => setVisible(false)}
        />
      </div>
      {renderClientSecrets()}
      <div className={style.logoLine}>
        <span className={style.logoLabel}>应用程序图标</span>
        <Logo />
      </div>
      <UpdateForm />
    </div>
  );
};

export default Detail;
