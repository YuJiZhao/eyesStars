// 固定配置信息
export default {
    appId: 251648,                                  // 应用id
    storageKey: {                                   // 本地存储key
        sToken: "sAuth",                            // 短token
        lToken: "lAuth",                            // 长token
        ids: "ids",                                 // 已读星星id加密列表
        stars: "stars",                             // 未读星星加密缓存
    },
    successCode: 200,                               // 接口成功状态码
    failCode: 400,                                  // 接口失败状态码
    layoutModeThreshold: 0.8,                       // 翻转阈值
    userInfoDefault: {                              // 用户信息默认占位值
        username: "未登录",
        email: "********@***.***"
    },
    userCenter: {                                   // 用户中心信息
        info: "http://user.eyescode.top",           // 信息页地址
        auth: "http://user.eyescode.top/OAuth2",    // 授权页地址
        redirectUrl: "/auth",                       // 重定向相对地址
    },
    tipDuration: 3000,                              // tip弹窗停留时间
    dialogDuration: 300,                            // 弹窗动画时间(tip除外)
    anonymousName: "匿名用户",                       // 用户匿名时名称展示
    dynamicZLevel: 0,                               // 动态元素层级
    staticZLevel: 1,                                // 静态元素层级
    moonSize: 70,                                   // 月亮尺寸
    sloganSizeNormal: 20,                           // 正常模式下的标语尺寸
    sloganSizeFlip: 16,                             // 翻转模式下的标语尺寸
    meteorStarWidth: 4,                             // 流星宽度
    meteorStarHeight: 100,                          // 流星高度
    meteorStarRotation: 30,                         // 流星基础倾斜角度
    meteorStarStartPosition: {                      // 流星出现范围(设备长度百分比)
        min: 0.4,
        max: 0.9
    },
    meteorStarAnimateDuration: 2000,                // 流星动画执行时间
    meteorStarChance: 0.05,                         // 普通流星出现概率(每个动画周期)
    decorateStarNum: 1 / 4000,                      // 背景星星密度(a/b，每b平方像素内a个)
    decorateStarN: {                                // 背景星星边数范围
        min: 6,
        max: 8
    },
    decorateStarR: {                                // 背景星星外半径范围
        min: 2,
        max: 3
    },
    decorateStarR0: {                               // 背景星星内半径范围
        min: 1,
        max: 1.5
    },
    decorateStarFill: {                             // 背景星星填充色范围
        min: 225,
        max: 255
    },
    decorateStarOpacity: {                          // 背景星星透明度范围
        min: 0.4,
        max: 0.8
    },
    messageStarN: {                                 // 寄语星星边数范围  
        min: 4,
        max: 4
    },
    messageStarR: {                                 // 寄语星星外半径范围
        min: 5,
        max: 7
    },
    messageStarR0: {                                // 寄语星星内半径范围
        min: 0.5,
        max: 2
    },
    messageStarFill: {                              // 寄语星星填充色范围
        min: 215,
        max: 255
    },
    messageStarOpacity: {                           // 寄语星星透明度范围
        min: 0.7,
        max: 1
    },
    messageStarNum: 1 / 10000,                      // 寄语星星密度(a/b，每b平方像素内a个)
    messageStarAnimateDuration: 5000,               // 寄语星星动画执行时间
}