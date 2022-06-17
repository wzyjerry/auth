import { useSelector } from 'umi';
import type { User } from 'umi';
import { Avatar } from 'antd';

const Navbar: React.FC = () => {
  const avatar = useSelector(({ user }: { user: User }) => user.avatar);
  return <Avatar size="large" src={avatar} />;
};

export default Navbar;
