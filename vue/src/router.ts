import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import ConvertExcelToTsv from './views/ExcelToTsv.vue';
Vue.use(Router);

export default new Router({
//  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      children: [
        {
          path: 'home',
          components: {
            main: Home,
          },
        },
        {
          path: 'convert_excel_to_tsv',
          components: {
            main: ConvertExcelToTsv,
          },
        },
      ],
    },
  ],
});
