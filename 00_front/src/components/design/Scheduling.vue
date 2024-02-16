<template>
  <div>
    <div class="scheduling">
      <v-radio-group v-model="scheValue">
        <v-radio label="平日毎日" value=0></v-radio>
        <v-radio label="1週間毎(毎週)" value=1></v-radio>
        <v-radio label="2週間毎(隔週)" value=2></v-radio>
        <v-radio label="3週間毎" value=3></v-radio>
        <v-radio label="4週間毎(毎月)" value=4></v-radio>
        <v-radio label="不定期" value=5></v-radio>
      </v-radio-group>
    </div>
    <div class="trigger">
      <div>
        <v-radio-group label="何曜日に設定しますか？" v-model="dayValue" inline v-if="scheValue >= 1 && scheValue <= 3">
          <!-- 日曜日を0スタートにすること -->
          <v-radio label="月曜" value=1></v-radio>
          <v-radio label="火曜" value=2></v-radio>
          <v-radio label="水曜" value=3></v-radio>
          <v-radio label="木曜" value=4></v-radio>
          <v-radio label="金曜" value=5></v-radio>
        </v-radio-group>
        <v-radio-group v-model="dateOrDay" inline v-if="scheValue == 4">
          <v-radio label="日付指定" value=0></v-radio>
          <v-radio label="曜日指定" value=1></v-radio>
        </v-radio-group>
      </div>
      <div>
        <v-select label="日付選択" :items="items" v-if="dateOrDay == 0" v-model="dateValue"></v-select>
        <v-radio-group label="何週目に設定しますか？" v-model="weekValue" inline v-if="dateOrDay == 1">
          <v-radio label="第1週" value=1></v-radio>
          <v-radio label="第2週" value=2></v-radio>
          <v-radio label="第3週" value=3></v-radio>
          <v-radio label="第4週" value=4></v-radio>
          <v-radio label="第5週" value=5></v-radio>
        </v-radio-group>
        <v-radio-group label="何曜日に設定しますか？" v-model="dayValue" inline v-if="dateOrDay == 1">
          <!-- 日曜日を0スタートにすること -->
          <v-radio label="月曜" value=1></v-radio>
          <v-radio label="火曜" value=2></v-radio>
          <v-radio label="水曜" value=3></v-radio>
          <v-radio label="木曜" value=4></v-radio>
          <v-radio label="金曜" value=5></v-radio>
      </v-radio-group>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Scheduling',
  data(){
    return {
      scheValue: 1,
      dateValue: 0,
      dayValue: 1,
      weekValue: 0,
      dateOrDay: -1,
      items: Array.from(Array(31), (_, i) => i + 1),
    }
  },
  methods:{
    GetScheduling(){
      return `${this.scheValue}-${this.weekValue}-${this.dayValue}-${this.dateValue}`
    },
  }
}
</script>

<style>
.scheduling {
  float: left;
  width: 180px;
  border-right: 1px solid rgb(204, 204, 204);
  box-sizing: border-box;
}

.trigger {
  display: inline-block;
  width: calc(100% - 180px); /* schedulingとdateを半分に分割し、左右の余白を除去 */
  box-sizing: border-box;
  padding: 0px 20px;
  vertical-align: top; /* 上端に揃える */
}

.v-select {
  width: 30%;
}
</style>