<template>
  <div class="setting">
    <!-- 子コンポーネントの関数をrefで呼ぶための記述 -->
    <person-manager ref="RefPersonManager"></person-manager>
    
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
          <v-tab value="option-5" v-if="show == true">
            <v-icon start>
              mdi-swap-horizontal
            </v-icon>
            順番変更
          </v-tab>
          <v-tab value="option-6" v-if="show == true">
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
              <select
                title="メンバーを選択してください"
                class="multi"
                multiple="multiple"
                size="20"
                >
                <option v-for="name in selectable" :key="name">{{ name }}</option>
              </select>
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
              ※注意※　更新すると元の順番には戻せません。<br>
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
    <v-btn flat rounded="pill" color="primary" @click="PageBack()">戻る</v-btn>
    <v-btn flat rounded="pill" color="primary" @click="Update(tab)" v-if="show == true">更新</v-btn>
    <setting-register v-if="show == false" @CbkRegister="CbkRegister"/>
  </div>
</template>

<script>
// import jQuery from 'jquery' //インポートししなくていいっぽい
import PersonManager from '@/components/model/PersonManager'
import SettingRegister from '@/components/view/SettingRegister'
import MailSetting from '@/components/view/VTabMailSetting'
// import MessageSetting from '@/components/view/VTabMessageSetting'
import ScheduleSetting from '@/components/view/VTabScheduleSetting'
import OwnerSetting from '@/components/view/VTabOwnerSetting'
import DragDropSwap from '@/components/design/DragDropSwap'
import InputTextField from '@/components/design/InputTextField'
import InputTextArea from '@/components/design/InputTextArea'

export default {
  name: 'Setting',
  components: {
    PersonManager,
    SettingRegister,
    MailSetting,
    // MessageSetting,
    ScheduleSetting,
    OwnerSetting,
    DragDropSwap,
    InputTextField,
    InputTextArea,
  },
  data(){
    return {
      // [Issue]本当はoption-1にしたいが、1にするとmultiselectがレンダリングされない
      tab: 'option-2',
      selectable: [],
      employeeNoList: [],
      toubanInfo: ["ID", "name", "message", "owner", "deleteKey"],
      member: ["toubanID", ["A", "B", "C", "D"]],
      show: true,
      hint: "当番名を設定してください",
    }
  },
  methods: {
    PageBack(){
      this.$router.push("/home")
    },
    CbkRegister(param){
      // 社員番号の一致チェック
      for(var i=0; i<this.employeeNoList.length; i++){
        if(this.employeeNoList[i] == param["owner"]){
          // オーナーと削除キーの格納
          this.toubanInfo["owner"] = param["owner"]
          this.toubanInfo["deleteKey"] = param["deleteKey"]
          // multiselectedから選択済み要素を格納
          this.StoreSelected()
          
          // リダイレクト
          this.$router.push("/home")
          return true //$emitはリターンを扱わないから飾り。とはいえ、breakと違って処理が終了する
        }
      }
      alert("データベースに存在しない社員番号です。")
    },
    // multiselectedから選択済み要素を格納
    StoreSelected(){
      const selected = document.getElementById("m-selected")
      var array = []
      for ( var i=0,l=selected.length; l>i; i++ ) {
        array.push(selected[i].value)
      }
      this.toubanInfo.member = array
    },
    Register(){
      // Update時の1つ1つの処理を全て実施すればいい
      Update("option-1")
      Update("option-2")
      Update("option-3")
      Update("option-4")
      Update("option-7")
    },
    Update(tab){
      var ret
      switch(tab){
        // 当番名
        case "option-1":
          this.toubanInfo.name = this.$refs.RefInputField.inputData
          break
        // メンバー編集
        case "option-2":
          this.StoreSelected()
          break
          // スケジュール設定
        case "option-3":
          var interval = this.$refs.RefSchedule.interval
          var startDate = this.$refs.RefSchedule.date
          break
        // メッセージ設定
        case "option-4":
          this.toubanInfo.message = this.$refs.RefInputArea.inputData
          break
          // 順番変更
        case "option-5":
          console.log(this.member)
          this.member = this.$refs.RefDragDropSwap.dataList.concat()
          console.log(this.member)
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
            alert("変更先オーナーもしくは削除キーの未設定です。")
          }
          break
        // メール配信設定
        case "option-7":
          var data = this.$refs.RefMail.GetMailSetting()
          break
      }
    }
  },
  mounted(){
    // URLによるタブの可視化制御のための処理
    const urlpath = this.$route.path
    if(urlpath == "/create"){
      this.show = false
    }
    // アドレス帳からデータ取得
    this.selectable = this.$refs.RefPersonManager.Members()
    this.employeeNoList = this.$refs.RefPersonManager.EmployeeNo()
    
    // DOMの更新を待ってから処理。nextTickがないとメンバーが表示されない
    this.$nextTick(function() {
      // MultiSelectedTableのマウント
      $('.multi').multiselectable({
        selectableLabel: '選択可能一覧',
        selectedLabel: '選択済み',
      });
    });
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