<template>
  <transition name="starAnimate">
    <div :class="'star star_' + starClass" v-if="show">
      <div class="close"><span @click="close">×</span></div>
      <div class="userInfo">
        <div class="avatar" :style="{backgroundImage: 'url(' + starData.avatar + ')'}"
        ></div>
        <div class="infoBox">
          <div class="name">{{ starData.name }}</div>
          <div class="time">{{ starData.createTime }}</div>
        </div>
      </div>
      <BaseDivider />
      <div class="content">{{ starData.content }}</div>
      <div class="footer">
        <div class="numbers">星序: {{ starData.id }}</div>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { LayoutModeEnum } from "@/constant/enum";
import { getStars } from "@/request/api";
import config from "@/config";
import { showTip } from "~~/bussiness/process";
import { aesDecrypt, aesEncrypt } from "@/utils/help";

let props = defineProps({
  show: { type: Boolean },
});

const monitorStore = useMonitor();
const processStore = useProcess();
const contextStore = useContext();
const starStore = useStar();
let starClass = ref(
  monitorStore.value.layoutMode == LayoutModeEnum.NORMAL ? "normal" : "flip"
);
let starData = reactive<{
  id: null | number;
  name: string;
  avatar: string;
  createTime: string;
  content: string;
}>({
  id: null,
  name: "",
  avatar: "",
  createTime: "",
  content: ""
})

function close() {
  processStore.value.isShowStar = false;
}

// 展示星星
function showStar() {
  let encryptStars = localStorage.getItem(config.storageKey.stars);
  if (starStore.value.stars.length != 0) {  // 启用内存数据
    let current = starStore.value.stars.pop();
    initStar(current!.Id, current!.Name, current!.Avatar, current!.CreateTime, current!.Content);
    updateLocalStorage();
  } else if(encryptStars != null) {  // 启用本地存储
    try {
      let current = updateLocalStorage();
      initStar(current!.Id, current!.Name, current!.Avatar, current!.CreateTime, current!.Content);
    } catch {
      localStorage.removeItem(config.storageKey.stars);
      showTip("没事不要动缓存啊兄弟");
    }
  } else {  // 获取新数据
    searchStars();
  }
}

// 给当前星星赋值
function initStar(id: number, name: string, avatar: string, createTime: string, content: string) {
  starData.id = id;
  starData.name = name;
  starData.avatar = avatar;
  starData.createTime = createTime;
  starData.content = content;
}

// 更新本地存储
function updateLocalStorage() {
  try {
    let encryptStars = localStorage.getItem(config.storageKey.stars);
    let stars = JSON.parse(aesDecrypt(encryptStars!, contextStore.value.aesKey!));
    let star = stars.pop();
    if (stars.length == 0) {
      localStorage.removeItem(config.storageKey.stars);
    } else {
      localStorage.setItem(config.storageKey.stars, aesEncrypt(JSON.stringify(stars), contextStore.value.aesKey!));
    }
    return star;
  } catch {
    localStorage.removeItem(config.storageKey.stars);
    showTip("没事不要动缓存啊兄弟");
    return {};
  }
}

// 获取星星
async function searchStars() {
  await getStars({
    ids: localStorage.getItem(config.storageKey.ids) || ""
  }).then(({code, msg, data}) => {
    if (code == config.successCode) {
      let jsonStars = "";
      try {
        jsonStars = aesDecrypt(data.stars, contextStore.value.aesKey!)
      } catch {  // 返回值不符合AES格式
        showTip("我说服务端错乱了你信不信");
      }
      if (jsonStars == "") {  // 密钥为空或错误
        showTip("我说服务端错乱了我还让它上线了你信不信");
      }
      // 展示星星
      let objectStars = JSON.parse(jsonStars);
      let current = objectStars.pop();
      initStar(current!.Id, current!.Name, current!.Avatar, current!.CreateTime, current!.Content);

      // 存入内存与客户端
      // 获取星星调用资源较多，因此需要批量获取，减少请求次数
      // 存入客户端是为了下次进入网站可以获取缓存数据，同时也能有效避免对同一用户返回已读数据
      // 存入内存是为了防止大佬手动改本地缓存造成请求次数增加，因此需要优先使用内存数据
      localStorage.setItem(config.storageKey.ids, data.ids);
      localStorage.setItem(config.storageKey.stars, aesEncrypt(JSON.stringify(objectStars), contextStore.value.aesKey!));
      starStore.value.ids = data.ids;
      starStore.value.stars = objectStars;
    } else {
      showTip(`获取星星失败!${msg}`);
    }
  })
}

watch(
  () => props.show,
  (value) => {
    if (value) {
      showStar();
    }
  }
);
</script>

<style scoped lang="scss">
@import "@/assets/scss/global.scss";

$dialogWidthNormal: 500px;
$dialogHeightNormal: 290px;
$dialogWidthFlip: 450px;
$dialogHeightFlip: 300px;
$avatarSize: 50px;
$nameHeight: 30px;

.starAnimate-enter-active {
  animation: zoomIn 0.3s;
}
.starAnimate-leave-active {
  animation: zoomOut 0.3s;
}

.star {
  position: fixed;
  z-index: $zIndexStar;
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
    }
    .infoBox {
      height: $avatarSize;
      margin-left: 10px;
      .name {
        height: $nameHeight;
        line-height: $nameHeight;
        font-size: 18px;
      }
      .time {
        line-height: $avatarSize - $nameHeight;
        font-size: 14px;
      }
    }
  }
  .content {
    margin-top: 10px;
  }
  .footer {
    display: flex;
    position: absolute;
    bottom: 10px;
    right: 10px;
    div {
      margin-left: 10px;
    }
  }
}

.star_normal {
  top: calc(45% - $dialogHeightNormal / 2);
  left: calc(50% - $dialogWidthNormal / 2);
  width: $dialogWidthNormal;
  height: $dialogHeightNormal;
}

.star_flip {
  top: calc(50% - $dialogHeightFlip / 2);
  left: calc(50% - $dialogWidthFlip / 2);
  transform: rotateZ(90deg);
  width: $dialogWidthFlip;
  height: $dialogHeightFlip;
}
</style>