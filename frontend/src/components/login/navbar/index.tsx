import { useSelector } from 'umi';
import type { User } from 'umi';
import { Avatar } from 'antd';

const Navbar: React.FC = () => {
  const user = useSelector(({ user }: { user: User }) => user);
  return (
    <Avatar size="large" src={user.avatar}>
      {user.nickname[0]}
    </Avatar>
  );
};

export default Navbar;
