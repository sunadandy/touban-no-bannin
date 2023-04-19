<template>
  <div class="vtab-setting">
    <div>
      <p>実施間隔と開始日を設定してください。<b>(デフォルトは今日から毎週)</b></p>
      <p>※不定期を設定した場合、メール配信は開始日に対する一度のみです。</p>
    </div>
    <v-radio-group v-model="interval" inline>
      <v-radio label="毎日" value="1"></v-radio>
      <v-radio label="毎週" value="2"></v-radio>
      <v-radio label="毎月" value="3"></v-radio>
      <v-radio label="不定期" value="4"></v-radio>
    </v-radio-group>
    <!-- 
      format:画面に表示する日時の表示形式
      model-type:v-modelの形式
      enable-time-picker:カレンダーに時間を表示するかどうか
      auto-apply:自動的にカレンダーを閉じるようにする(日付クリックが選択となる)
      inline:カレンダーを常に表示
    -->
    <VueDatePicker
      v-model="date"
      :format="format"
      model-type="yyyy-MM-dd"
      :enable-time-picker=false
      auto-apply
      inline
    ></VueDatePicker>
    {{ date }}
  </div>
</template>

<script>
import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'

export default {
  components: { VueDatePicker },
  name: "ScheduleSetting",
  data() {
    var _today = new Date();
    var year = _today.getFullYear()
    var month = _today.getMonth() + 1
    var day = _today.getDate()
    var today = year + "/" + month + "/" + day
    return {
      interval: 2,
      date: today,
      format: function(date){
        const day = date.getDate();
        const month = date.getMonth() + 1;
        const year = date.getFullYear();
        return `${year}/${month}/${day}`;
      }
    };
  },
}
</script>

<style scoped>
.setting{
  padding-top: 1em;
}
.date-picker {
  display: none;
}
</style>