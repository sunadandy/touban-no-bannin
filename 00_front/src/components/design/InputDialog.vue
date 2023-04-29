<template>
  <div>
    <v-card>
      <v-card-text>
        <v-container>
          <v-row>
            <!-- 親から受け取ったヒントの数だけ入力フォームを設置 -->
            <div v-for="(hint, index) in hints" :key="index">
              <v-col cols="12">
                <input-text-field :ref="`RefInputField${index}`" v-bind:hint='hint'/>
              </v-col>
            </div>
            <v-col cols="12" sm="6">
            </v-col>
          </v-row>
        </v-container>
        <!-- <small>*登録成功時、自動でホームに戻ります。</small> -->
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue-darken-1" variant="text" @click="Close">閉じる</v-btn>
        <v-btn color="blue-darken-1" variant="text" @click="Submit">送信</v-btn>
      </v-card-actions>
    </v-card>
    <overlay ref="RefOverlay" v-bind:waitTime="waitTime"/>
  </div>
</template>

<script>
import InputTextField from '@/components/design/InputTextField'
import Overlay from '@/components/design/Overlay'

export default {
  name: 'InputDialog',
  components: { 
    InputTextField,
    Overlay,
  },
  data(){
    return {
      waitTime: 500,    //オーバーレイコンポーネントに渡すパラメータ
      returnData: [],
    }
  },
  props:{
    hints: {
      type: Array,
      required: true,
      default: () => [],
    },
  },
  methods: {
    Close(){
      this.$emit('clickClose')
    },
    Submit(){
      this.returnData = []  //初期化。これがないとずっとpushされる
      this.hints.forEach((hint, index) => {
        this.returnData.push(this.$refs[`RefInputField${index}`][0].inputData)
      });
      // 演出
      this.$refs.RefOverlay.SetOverlay()
      setTimeout(() => {
        this.$emit('clickSubmit', this.returnData)
      }, this.waitTime)
    },
  },
}
</script>