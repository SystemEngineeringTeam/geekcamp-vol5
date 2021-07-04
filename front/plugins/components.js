import Vue from 'vue'

// componentsファイルにあるグローバルにしたいコンポーネントをimport
import Celebration from '~/components/Celebration.vue'
// それを今回は'Button'というコンポーネント名で設定。
Vue.component('Celebration', Celebration)
