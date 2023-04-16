<template>
  <div class="setting">
    <div>
        メンバーに伝えたいメッセージを入力してください。(サーバの参照資料等)<br>
    </div>
    <input-text-area v-bind:hint="hint"/>
  </div>
</template>

<script>
import InputTextArea from '@/components/design/InputTextArea.vue'

export default {
  name: 'MassageSetting',
  components: { InputTextArea },
  data(){
    return{
      hint: ""
    }
  }
}
</script>

<style scoped>
.setting{
  padding-top: 1em;
}
</style>