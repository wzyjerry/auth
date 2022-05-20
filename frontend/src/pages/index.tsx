import { Reject, Throw } from '@/util';
import { Token } from '@/util/token';
import styles from './index.less';

const App = () => {
  const helper = new Token();
  var result = helper.Load();
  if (result instanceof Reject) {
    console.log(result.error.name);
  } else {
    console.log(result.val);
  }

  var err = helper.Save('token');
  if (err instanceof Reject) {
    console.log(err.error.name);
  }

  var result = helper.Load();
  if (result instanceof Reject) {
    console.log(result.error.name);
  } else {
    console.log(result.val);
  }

  return (
    <div>
      <h1 className={styles.title}>Auth frontend</h1>
    </div>
  );
};

export default App;
