import { defineConfig } from 'umi';

export default defineConfig({
  dva: {
    immer: true,
    hmr: true,
  },
  nodeModulesTransform: {
    type: 'none',
  },
  fastRefresh: {},
});
