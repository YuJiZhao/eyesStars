<template>
  <transition name="publishAnimate">
    <div :class="'publish publish_' + publishClass" v-if="show">
      <div class="close"><span @click="close">×</span></div>
      <!-- 个人信息 -->
      <div class="userInfo">
        <div class="avatar" 
          :style="{ backgroundImage: 'url(' + userStore.avatar + ')' }"
          @click="goUserCenter"
          ></div>
        <div class="infoBox">
          <div class="name">{{ userStore.username }}</div>
          <div class="email">{{ userStore.email }}</div>
        </div>
      </div>
      <!-- 发布区 -->
      <div class="publishBox">
        <textarea v-model="body.content" :placeholder="placeholder" :maxlength="msgMaxLen"></textarea>
      </div>
      <!-- 功能区 -->
      <div class="ribbon">
        <input class="alias" type="text" placeholder="使用化名" v-model="body.name"/>
        <div
          :class="{ btn: true, anonymousBox: true, active: body.anonymous }"
          @click="body.anonymous = !body.anonymous"
        >
          是否匿名
        </div>
        <div class="btn publishBtn" @click="publishStar">点击发布</div>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { DialogEnum, LayoutModeEnum } from "@/constant/enum";
import { addStar, loginTrack } from "@/request/api";
import { showDialog, showTip } from "@/bussiness/process";
import config from "@/config";
import { aesDecrypt, aesEncrypt, browser, OS } from "@/utils/help";
import { LoginPatternEnum } from "@/constant/enum";

defineProps({
  show: { type: Boolean },
});

const route = useRoute();
const runtimeConfig = useRuntimeConfig();
const monitorStore = useMonitor();
const processStore = useProcess();
const contextStore = useContext();
const userStore = useUser();
const starStore = useStar();
let publishClass = ref(
  monitorStore.value.layoutMode == LayoutModeEnum.NORMAL ? "normal" : "flip"
);
let placeholder = ref(
  `仅登录用户可发布，最多${contextStore.value.messageMaxLen}字。`
);
let msgMaxLen = ref(contextStore.value.messageMaxLen);
let body = reactive({
  content: "",
  name: "",
  anonymous: false,
});
let lock = false;  // 锁，防止多次提交

function close() {
  processStore.value.isShowPublish = false;
}

function jumpTo(site: string) {
  window.open(site);
}

// 未登录则前往登录，已登录则查看信息
function goUserCenter() {
  let pattern;
  if (userStore.value.isLoggedin) {
    pattern = LoginPatternEnum.INFO;
    window.open(`${config.userCenter.info}?clientId=${config.appId}`);
  } else {
    pattern = LoginPatternEnum.LOGIN;
    window.open(`${config.userCenter.auth}?clientId=${config.appId}&redirectUrl=${runtimeConfig.public.siteUrl + config.userCenter.redirectUrl}`);
  }
  let jsonParam = JSON.stringify({
    pattern,
    path: route.fullPath,
    os: OS(),
    browser: browser()
  });
  loginTrack({ package: encodeURIComponent(aesEncrypt(jsonParam, contextStore.value.aesKey!)) });
}

// 将发布的星星存入内存与本地存储
function storageNewStar(id: number, name: string, avatar: string, createTime: string, content: string) {
  let newStar = {
    Id: id,
    Name: name,
    Avatar: avatar,
    CreateTime: createTime,
    Content: content
  };
  starStore.value.stars.push(newStar);
  let encryptStars = localStorage.getItem(config.storageKey.stars);
  if (encryptStars != null) {
    try {
      let jsonStars = aesDecrypt(encryptStars, contextStore.value.aesKey!);
      let objectStars = JSON.parse(jsonStars);
      objectStars.push(newStar);
      localStorage.setItem(config.storageKey.stars, aesEncrypt(JSON.stringify(objectStars), contextStore.value.aesKey!));
    } catch {
      localStorage.setItem(config.storageKey.stars, aesEncrypt(JSON.stringify([newStar]), contextStore.value.aesKey!));
    }
  } else {
    localStorage.setItem(config.storageKey.stars, aesEncrypt(JSON.stringify([newStar]), contextStore.value.aesKey!));
  }
}

// 发布星星
async function publishStar() {
  if (!userStore.value.isLoggedin) {
    showTip("请先登录");
    return;
  }
  if (!lock) {
    if (body.content == "") {
      showTip("内容不能为空");
      return;
    }
    if (body.content.length > contextStore.value.messageMaxLen!) {
      showTip(`内容不能超过${contextStore.value.messageMaxLen}字`);
      return;
    }

    lock = true;
    await addStar({...body}).then(({code, msg, data}) => {
      if (code == config.successCode) {
        showTip(`第${data.id}颗星星发布成功`);
        
        // 跳转star弹窗
        storageNewStar(
          data.id,
          data.name == "" ? (body.anonymous ? config.anonymousName : userStore.value.username) : data.name,
          body.anonymous ? contextStore.value.defaultAvatar! : userStore.value.avatar,
          data.createTime,
          data.content
        );
        showDialog(DialogEnum.STAR);

        // 清除数据
        body.content = "";
        body.name = "";
        body.anonymous = false;
      } else {
        showTip(msg);
      }
      lock = false;
    });
  }
}
</script>

<style scoped lang="scss">
@import "@/assets/scss/global.scss";

$dialogWidthNormal: 500px;
$dialogHeightNormal: 290px;
$dialogWidthFlip: 450px;
$dialogHeightFlip: 300px;
$avatarSize: 50px;
$nameHeight: 30px;

.publishAnimate-enter-active {
  animation: zoomIn 0.3s;
}
.publishAnimate-leave-active {
  animation: zoomOut 0.3s;
}

.publish {
  position: fixed;
  z-index: $zIndexPublish;
  background: rgba(#fff, 0.8);
  border-radius: 10px;
  box-shadow: 0 0 5px #fff;
  -webkit-box-shadow: 0 0 5px #fff;
  -moz-box-shadow: 0 0 5px #fff;
  padding: 5px 10px;
  color: $normalColor;
  .close {
    text-align: right;
    span {
      cursor: pointer;
      transition: 0.3s;
      &:hover {
        color: $activeColor;
      }
    }
  }
  .userInfo {
    display: flex;
    justify-content: flex-start;
    .avatar {
      width: $avatarSize;
      height: $avatarSize;
      background-size: 100% 100%;
      border-radius: 50%;
      cursor: pointer;
    }
    .infoBox {
      height: $avatarSize;
      margin-left: 10px;
      .name {
        height: $nameHeight;
        line-height: $nameHeight;
        font-size: 18px;
      }
      .email {
        line-height: $avatarSize - $nameHeight;
        font-size: 14px;
      }
    }
  }
  .publishBox {
    margin-top: 10px;
    textarea {
      width: calc(100% - 20px);
      height: 130px;
      background: none;
      border-radius: 5px;
      border: 2px solid #999;
      resize: none;
      padding: 10px;
      color: $normalColor;
      outline: none;
    }
  }
  .ribbon {
    height: 30px;
    margin-top: 15px;
    display: flex;
    justify-content: flex-end;
    .alias {
      width: 80px;
      height: 30px;
      border-radius: 5px;
      border: 2px solid #999;
      text-align: center;
      color: $normalColor;
      cursor: pointer;
      background: none;
      outline: none;
      font-size: 16px;
      transition: 0.3s;
      &::-webkit-input-placeholder {
        color: $normalColor;
        font-size: 15px;
        transition: 0.3s;
        transform: translateY(1.5px);
      }
      &:-moz-placeholder {
        color: $normalColor;
        opacity: 1;
        font-size: 15px;
        transition: 0.3s;
        transform: translateY(1.5px);
      }
      &::-moz-placeholder {
        color: $normalColor;
        opacity: 1;
        font-size: 15px;
        transition: 0.3s;
        transform: translateY(1.5px);
      }
      &:-ms-input-placeholder {
        color: $normalColor;
        font-size: 15px;
        transition: 0.3s;
        transform: translateY(1.5px);
      }
      &:hover {
        background: #999;
        color: #fff;
        &::-webkit-input-placeholder {
          color: #fff;
        }
        &:-moz-placeholder {
          color: #fff;
        }
        &::-moz-placeholder {
          color: #fff;
        }
        &:-ms-input-placeholder {
          color: #fff;
        }
      }
      &:focus {
        background: none;
        width: 120px;
        color: $normalColor;
        &::-webkit-input-placeholder {
          color: #999;
        }
        &:-moz-placeholder {
          color: #999;
        }
        &::-moz-placeholder {
          color: #999;
        }
        &:-ms-input-placeholder {
          color: #999;
        }
      }
    }
    .btn {
      width: 80px;
      height: 30px;
      border-radius: 5px;
      border: 2px solid #999;
      margin-left: 10px;
      text-align: center;
      line-height: 30px;
      cursor: pointer;
      transition: 0.3s;
      &:hover {
        background: #999;
        color: #fff;
      }
    }
    .active {
      background: #999;
      color: #fff;
    }
  }
}

.publish_normal {
  top: calc(45% - $dialogHeightNormal / 2);
  left: calc(50% - $dialogWidthNormal / 2);
  width: $dialogWidthNormal;
  height: $dialogHeightNormal;
}

.publish_flip {
  top: calc(50% - $dialogHeightFlip / 2);
  left: calc(50% - $dialogWidthFlip / 2);
  transform: rotateZ(90deg);
  width: $dialogWidthFlip;
  height: $dialogHeightFlip;
}
</style>