<template>
  <div class="setting">
    <v-card class="setting-vcard">
      <v-toolbar
        color="primary"
      >
        <v-toolbar-title>当番の作成・編集</v-toolbar-title>
      </v-toolbar>
      <div class="d-flex flex-row">
        <v-tabs
          v-model="tab"
          direction="vertical"
          color="primary"
        >
          <v-tab value="option-1">
            <v-icon start>
              mdi-file-edit
            </v-icon>
            当番名
          </v-tab>
          <v-tab value="option-2">
            <v-icon start>
              mdi-account-search
            </v-icon>
            メンバー編集
          </v-tab>
          <v-tab value="option-3">
            <v-icon start>
              mdi-clock-edit
            </v-icon>
            スケジューリング
          </v-tab>
          <v-tab value="option-4">
            <v-icon start>
              mdi-message-text
            </v-icon>
            メッセージ
          </v-tab>
          <v-tab value="option-5" v-if="limited == false">
            <v-icon start>
              mdi-swap-horizontal
            </v-icon>
            順番変更
          </v-tab>
          <v-tab value="option-6" v-if="limited == false">
            <v-icon start>
              mdi-account-switch
            </v-icon>
            オーナー変更
          </v-tab>
          <v-tab value="option-7">
            <v-icon start>
              mdi-email-arrow-right
            </v-icon>
            メール配信設定
          </v-tab>
        </v-tabs>
        <v-window v-model="tab">
          <v-window-item value="option-1">
            <v-card flat>
              <v-card-text>
                <input-text-field ref="RefInputField" v-bind:hint='hint'/>
              </v-card-text>
            </v-card>
          </v-window-item>
          <v-window-item value="option-2">
            <v-card flat>
              <member-select ref="RefMemberSelect" v-bind:memberInfo="memberInfo"/>
            </v-card>
          </v-window-item>
          <v-window-item value="option-3">
            <v-card flat>
              <v-card-text>
                <schedule-setting ref="RefSchedule"/>
              </v-card-text>
            </v-card>
          </v-window-item>
          <v-window-item value="option-4">
            <v-card flat>
              メンバーに伝えたいメッセージを入力してください。(サーバの参照資料等)<br>
              <input-text-area ref="RefInputArea"/>
            </v-card>
          </v-window-item>
          <v-window-item value="option-5">
            <v-card flat>
              ※注意※  更新すると元の順番には戻せません。<br>
              <!-- オリジナルを渡すとスワップ時にオリジナルまでスワップされるのでコピーを渡す -->
              <drag-drop-swap ref="RefDragDropSwap" v-bind:dataList="member.concat()"/>
            </v-card>
          </v-window-item>
          <v-window-item value="option-6">
            <v-card flat>
              <owner-setting v-bind:memberList="member" ref="RefOwner"></owner-setting>
            </v-card>
          </v-window-item>
          <v-window-item value="option-7">
            <v-card flat>
              <mail-setting ref="RefMail"/>
            </v-card>
          </v-window-item>
        </v-window>
      </div>
    </v-card>
    <v-btn flat rounded="pill" color="primary" @click="PageBack">戻る</v-btn>
  </div>
</template>

<script>
import MailSetting from '@/components/view/VTabMailSetting'
import ScheduleSetting from '@/components/view/VTabScheduleSetting'
import OwnerSetting from '@/components/view/VTabOwnerSetting'
import MemberSelect from '@/components/view/VTabMemberSelect'
import DragDropSwap from '@/components/design/DragDropSwap'
import InputTextField from '@/components/design/InputTextField'
import InputTextArea from '@/components/design/InputTextArea'

export default {
  name: 'Setting',
  components: {
    MailSetting,
    ScheduleSetting,
    OwnerSetting,
    MemberSelect,
    DragDropSwap,
    InputTextField,
    InputTextArea,
  },
  data(){
    return {
      // [Issue]本当はoption-1にしたいが、1にするとmultiselectがレンダリングされない
      tab: 'option-2',
      hint: "当番名を設定してください",
      memberInfo: [],   //{"employeeNo", "name", "email"}のJSON配列
    }
  },
  props:{
    // タブの表示非表示を制御するためのパラメータ
    limited:{
      type: Boolean,
      require: true,
      default: () => false,
    },
    current_memberInfo: {
      type: Array,
      require: true,
      default: () => [],
    }
  },
  methods: {
    PageBack(){
      this.$router.push("/home")
    },
    GetSelected(){
      return this.$refs.RefMemberSelect.GetSelected()
    },
    Validation(){
      if(this.limited == true){
        const result = this.ValidationAll()
        return result
      }else{
        const result = this.ValidationIndividual(this.tab)
        return result
      }
    },
    ValidationAll(){
      const validations  = [
        this.ValidationIndividual("option-1"),
        this.ValidationIndividual("option-2"),
        this.ValidationIndividual("option-3"),
        this.ValidationIndividual("option-4"),
        this.ValidationIndividual("option-7"),
      ]
      const isValid = validations.every(validation => validation.status);   // validation.statusが全てtrueならtrueを返す
      const dataArray = validations.map(validation => {
        return {
          data: validation.data
        }
      })
      const result = {status: isValid, data: dataArray}
      
      return result
    },
    ValidationIndividual(tab){
      var data
      var status
      switch(tab){
        // 当番名
        case "option-1":
          try{
            data = this.$refs.RefInputField.inputData
            status = true
          }catch{
            // 子コンポーネントがインスタンス化されていない状態でdataプロパティにアクセスした場合に例外発生
            data = null
            status = false
          }
          break
        // メンバー編集
        case "option-2":
          try{
            data = this.GetSelected()
            if(data.length == 0){
              status = false
            }else{
              status = true
            }
          }catch{
            // 子コンポーネントがインスタンス化されていない状態でdataプロパティにアクセスした場合に例外発生
            data = null
            status = false
          }
          break
        // スケジュール設定
        case "option-3":
          try{
            const interval = this.$refs.RefSchedule.interval
            const startDate = this.$refs.RefSchedule.date
            data = {interval, startDate}  //連想配列の時点でintから勝手にStringになるっぽい
            status = true
          }catch{
            // 子コンポーネントがインスタンス化されていない状態でdataプロパティにアクセスした場合に例外発生
            data = null
            status = false
          }
          break
        // メッセージ設定
        case "option-4":
          try{
            data = this.$refs.RefInputArea.inputData
            status = true
          }catch{
            // 子コンポーネントがインスタンス化されていない状態でdataプロパティにアクセスした場合に例外発生
            data = null
            status = false
          }
          break
          // 順番変更
        case "option-5":
          try{
            data = this.$refs.RefDragDropSwap.dataList.concat()
            status = true
          }catch{
            // 子コンポーネントがインスタンス化されていない状態でdataプロパティにアクセスした場合に例外発生
            data = null
            status = false
          }
          break
        // オーナー変更
        case "option-6":
          var data = this.$refs.RefOwner.GetData()
          var nextOwner = data.nextOwner
          var deleteKey = data.newKey
          if(nextOwner != null && deleteKey != null){
            this.toubanInfo.owner = nextOwner
            this.toubanInfo.deleteKey = deleteKey
          }else{
            alert("変更先オーナーもしくはパスワードの未設定です。")
          }
          break
        // メール配信設定
        case "option-7":
          try{
            data = this.$refs.RefMail.GetMailSetting()
            status = true
          }catch{
            // 子コンポーネントがインスタンス化されていない状態でdataプロパティにアクセスした場合に例外発生
            data = null
            status = false
          }
          break
      }

      return {
        tab: tab,
        data: data,
        status: status, //現状、true固定。falseが返るケースが発生した際は修正
      }
    }
  },
  created(){
    this.memberInfo = this.current_memberInfo.map((info) => ({
      name: info.name,
      emplyoeeNo: info.emplyoee_number,
      email: info.email,
    }))
  },
}
</script>

<style class='foobar'>
.multiselectable {
  justify-content: left;
  width: auto;
  display: flex;
  overflow: hidden;
  width: 200%;
  margin-bottom: 10px;
}

.multiselectable select {
  float: left;
  height: 500px;
  width: 400px;
  border-width: 1px;
  border-style: solid;
  border-color:slategrey;
}
.multiselectable div {
  float: left;
  height: auto;
  width: 400px;
  font-weight: bold;
  color: slategray;
}

.multiselectable div * {
  display: block;
  margin: 0 auto;
}

.multiselectable div {
  display: inline;
}

.multiselectable .m-selectable-controls {
  margin-top: 3em;
  width: 50px;
}

.multiselectable .m-selectable-controls button {
    margin-top: 1em;
    padding: .1em .3em;
    background-color: slategray;
    color: white;
}
.v-text-field{
  width: 500px;
  margin: auto;
  padding-top: 1em;
}
.multi select{
  display: block;
}
.v-tabs {
  width: 200px;
}
.v-window{
  padding-left: 1em;
  padding-top: 1em;
}
.vtab-setting{
  margin-left: 10px;
}
</style>