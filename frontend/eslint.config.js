import { defineFlatConfig } from 'eslint-define-config';
import vue from 'eslint-plugin-vue';
import vueParser from 'vue-eslint-parser';

export default defineFlatConfig([
  {
    files: ['**/*.vue'],
    languageOptions: {
      parser: vueParser,
      parserOptions: {
        parser: '@babel/eslint-parser',
        requireConfigFile: false, // no need for babel config file
        ecmaVersion: 2021,
        sourceType: 'module',
      },
    },
    plugins: {
      vue,
    },
    rules: {
      ...vue.configs['vue3-recommended'].rules,
      'vue/multi-word-component-names': 'off',
    },
  },
]);
