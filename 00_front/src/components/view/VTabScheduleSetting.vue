<template>
  <div class="vtab-setting">
    実施間隔と開始日を設定してください。(デフォルトは今日から毎週)<br>
    ※不定期を設定した場合、メール配信は初回実施日に対する一度のみです。<br>
    <div class="vtab-scheduling">
      <scheduling ref="RefScheduling"/>
    </div>
    <!-- 
      format:画面に表示する日時の表示形式
      model-type:v-modelの形式
      enable-time-picker:カレンダーに時間を表示するかどうか
      auto-apply:自動的にカレンダーを閉じるようにする(日付クリックが選択となる)
      inline:カレンダーを常に表示
    -->
    <div class="calendar">
      <VueDatePicker
        v-model="date"
        :format="format"
        model-type="yyyy-MM-dd"
        :enable-time-picker=false
        auto-apply
        inline
      ></VueDatePicker>
    </div>
    開始日：{{ date }}
  </div>
</template>

<script>
import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'
import Scheduling from '@/components/design/Scheduling'
import {format, startOfToday} from 'date-fns'

export default {
  components: {
    VueDatePicker,
    Scheduling,
  },
  name: "ScheduleSetting",
  data() {
    return {
      date: format(startOfToday(), 'yyyy-MM-dd')
    };
  },
  methods:{
    GetScheduleSetting(){
      return {
        scheduling: this.$refs.RefScheduling.GetScheduling(),
        startDate: this.date,
      }
    }
  }
}
</script>

<style>
.vtab-scheduling {
  display: flex;
  margin: 1rem;
}
.calendar {
  display: flex;
  justify-content: center;
}
</style>