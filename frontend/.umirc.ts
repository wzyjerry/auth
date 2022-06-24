import { defineConfig } from 'umi';
import theme from './theme';
export default defineConfig({
  dva: {
    immer: true,
    hmr: true,
  },
  nodeModulesTransform: {
    type: 'none',
  },
  fastRefresh: {},
  // mfsu: {},
  theme,
});
