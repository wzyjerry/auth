// .eslintrc.js
module.exports = {
  extends: [require.resolve('@umijs/fabric/dist/eslint')],

  // 自定义全局变量
  globals: {
    REACT_APP_ENV: true,
    REACT_APP_API_URL: true,
  },
  // 根据个人习惯自定义rule
  rules: {
    'no-console': 'off',
    'react-hooks/exhaustive-deps': 'error', // react-hooks 依赖检查
    'no-empty': 'off', // catch{} 允许为空
    '@typescript-eslint/no-shadow': ['off'], // 当前作用域变量名不能与父级作用域变量同名
  },
};
