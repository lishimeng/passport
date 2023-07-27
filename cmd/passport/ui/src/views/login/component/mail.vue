<template>
  <el-form size="large" class="login-content-form">
    <el-form-item class="login-animation1">
      <el-input text :placeholder="$t('message.mail.mail')" v-model="state.ruleForm.userName" clearable
                autocomplete="off">
        <template #prefix>
          <i class="iconfont icon-gerenzhongxin el-input__icon"></i>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item class="login-animation2">
      <el-col :span="15">
        <el-input text maxlength="4" :placeholder="$t('message.mail.code')" v-model="state.ruleForm.code" clearable
                  autocomplete="off">
          <template #prefix>
            <el-icon class="el-input__icon">
              <ele-Position/>
            </el-icon>
          </template>
        </el-input>
      </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="8">
        <el-button v-if="state.showCode" @click="sendCode" v-waves class="login-content-code">{{
            state.codeText
          }}
        </el-button>
        <el-button v-else v-waves class="login-content-code">{{ state.codeText }}</el-button>
      </el-col>
    </el-form-item>
    <el-form-item class="login-animation3">
      <el-button @keyup.enter="login" @click="login" round type="primary" v-waves class="login-content-submit">
        <span>{{ $t('message.mobile.btnText') }}</span>
      </el-button>
    </el-form-item>
    <div class="font12 mt30 login-animation4 login-msg">{{ $t('message.mobile.msgText') }}</div>
  </el-form>
</template>

<script setup lang="ts" name="loginMail">

import {computed, reactive} from "vue";
import {ElMessage} from "element-plus";
import {signInCodeApi} from "/@/api/login";
import {Local, Session} from "/@/utils/storage";
import {formatAxis} from "/@/utils/formatTime";
import {initFrontEndControlRoutes} from "/@/router/frontEnd";
import {initBackEndControlRoutes} from "/@/router/backEnd";
import {NextLoading} from "/@/utils/loading";
import {useThemeConfig} from "/@/stores/themeConfig";
import {storeToRefs} from "pinia";
import {useRoute, useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import {sendCodeApi} from "/@/api/send";
import {getOpenUrl} from "/@/utils/openUrl";

const storesThemeConfig = useThemeConfig();
const {themeConfig} = storeToRefs(storesThemeConfig);
const route = useRoute();
const router = useRouter();
// 定义变量内容
const {t} = useI18n();
const state = reactive({
  ruleForm: {
    userName: "",
    code: ""
  },
  showCode: true,
  codeTimeNum: 60,
  setIntervalTime: null,
  codeText: "获取验证码",
  loading: {
    signIn: false,
  },
  authCode:""
})
const login = () => {
  if (!state.ruleForm.userName) {
    ElMessage.error("邮箱不能为空！")
    return
  }
  if (!state.ruleForm.code) {
    ElMessage.error("验证码不能为空！")
    return
  }
  if(!checkMail(state.ruleForm.userName)){
    ElMessage.error("请输入正确的邮箱！")
    return
  }
  state.loading.signIn = true;
  signInCodeApi({
    userName: state.ruleForm.userName,
    code: state.ruleForm.code,
    codeLoginType: "mail"
  }).then(res => {
    if (res && res.code == 200) {
      // 存储 token 到浏览器缓存
      Local.set('token', res.token);
      Session.set('token', res.token);
      //外部登录
      var referrer=document.referrer
      var localHref=window.location.href
      var openUrl=getOpenUrl(referrer,Local.get("openPath"),res.token)
      if(referrer&&localHref.indexOf(referrer)<0){
        window.location.replace(openUrl)
      }else {
        signInSuccess()
      }
    } else {
      ElMessage.error("登录失败！")
      state.loading.signIn = false;
    }
  })
}
const checkMail=(val:any)=>{
  const regEmail = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/
  if(regEmail.test(val)){
     return true;
  }
  return false;
}
// 时间获取
const currentTime = computed(() => {
  return formatAxis(new Date());
});
// 登录成功后的跳转
const signInSuccess = async () => {
  let isNoPower = null
  if (!themeConfig.value.isRequestRoutes) {
    // 前端控制路由，2、请注意执行顺序
    isNoPower = await initFrontEndControlRoutes();
  } else {
    // 模拟后端控制路由，isRequestRoutes 为 true，则开启后端控制路由
    // 添加完动态路由，再进行 router 跳转，否则可能报错 No match found for location with path "/"
    isNoPower = await initBackEndControlRoutes();
  }
  if (isNoPower) {
    ElMessage.warning('抱歉，您没有登录权限');
    Session.clear();
  } else {
    // 初始化登录成功时间问候语
    let currentTimeInfo = currentTime.value;
    // 登录成功，跳到转首页
    // 如果是复制粘贴的路径，非首页/登录页，那么登录成功后重定向到对应的路径中
    if (route.query?.redirect) {
      router.push({
        path: <string>route.query?.redirect,
        query: Object.keys(<string>route.query?.params).length > 0 ? JSON.parse(<string>route.query?.params) : '',
      });
    } else {
      router.push('/');
    }
    // 登录成功提示
    const signInText = t('message.signInText');
    ElMessage.success(`${currentTimeInfo}，${signInText}`);
    // 添加 loading，防止第一次进入界面时出现短暂空白
    NextLoading.start();
  }
  state.loading.signIn = false;
}
const authCode = async () => {
  // state.authCode=(Math.round(Math.random()*(9999-1000)+1000)).toString()
  // console.log(state.authCode)
  await sendCodeApi({
    // code:state.authCode,
    codeLoginType:"mail",
    receiver:state.ruleForm.userName
  }).then(res=>{
    if(res&&res.code==200){
      ElMessage.success("验证码发送成功，请注意查收！")
    }else{
      ElMessage.error("验证码发送失败！")
      state.showCode = true
      state.codeText = "获取验证码";
      clearInterval(state.setIntervalTime);
    }
  })
}
const sendCode = () => {
  if (!state.ruleForm.userName) {
    ElMessage.error("邮箱不能为空！")
    return
  }
  if(!checkMail(state.ruleForm.userName)){
    ElMessage.error("请输入正确的邮箱！")
    return
  }
  state.showCode = false
  state.codeTimeNum = 60
  state.setIntervalTime = setInterval(countDown, 1000)
  authCode()
}
const countDown = () => {
  // console.log(state.codeTimeNum)
  state.codeTimeNum--;
  if (state.codeTimeNum <= 0) {
    state.showCode = true
    state.codeText = "获取验证码";
    clearInterval(state.setIntervalTime);
  } else {
    state.codeText = state.codeTimeNum + "s"
  }
}
</script>

<style scoped lang="scss">
.login-content-form {
  margin-top: 20px;
  @for $i from 1 through 4 {
    .login-animation#{$i} {
      opacity: 0;
      animation-name: error-num;
      animation-duration: 0.5s;
      animation-fill-mode: forwards;
      animation-delay: calc($i/10) + s;
    }
  }
  .login-content-code {
    width: 100%;
    padding: 0;
  }
  .login-content-submit {
    width: 100%;
    letter-spacing: 2px;
    font-weight: 300;
    margin-top: 15px;
  }
  .login-msg {
    color: var(--el-text-color-placeholder);
  }
}
</style>