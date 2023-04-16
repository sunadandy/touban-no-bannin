<template>
  <div class="vtab-setting">
    <div>
      <p>実施間隔と開始日を設定してください。</p>
      <p>※不定期を設定した場合、次回実施日を個別設定しない限り、メール通知は行われません。</p>
    </div>
    <v-radio-group inline>
      <v-radio label="毎日" value="1"></v-radio>
      <v-radio label="毎週" value="2"></v-radio>
      <v-radio label="毎月" value="3"></v-radio>
      <v-radio label="不定期" value="4"></v-radio>
    </v-radio-group>
    <VueDatePicker v-model="date" :format="format" inline></VueDatePicker>
  </div>
</template>

<script>
import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'

export default {
  components: { VueDatePicker },
  name: "ScheduleSetting",
  data() {
    return {
      date: null,
      format: (date) => {
        const day = date.getDate();
        const month = date.getMonth() + 1;
        const year = date.getFullYear();
        return `${year}/${month}/${day}`;
      }
    };
  }
}
</script>

<style scoped>
.setting{
  padding-top: 1em;
}
</style>