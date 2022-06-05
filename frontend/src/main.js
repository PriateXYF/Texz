import Vue from 'vue'
import App from './App.vue'
import { Button, Col, Row, Icon, Field, Popup, Search, Cell, CellGroup } from 'vant'

Vue.use(Button)
Vue.use(Button)
Vue.use(Col)
Vue.use(Row)
Vue.use(Icon)
Vue.use(Field)
Vue.use(CellGroup)
Vue.use(Popup)
Vue.use(Search)
Vue.use(Cell)
Vue.use(CellGroup)


Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
