// 固定配置信息
export default {
    apiBaseUrl: "http://127.0.0.1:8888",    // 接口服务器地址：TODO: 记得改成环境变量
    layoutModeThreshold: 0.8,               // 翻转阈值
    dynamicZLevel: 0,                      // 动态元素层级
    staticZLevel: 1,                       // 静态元素层级
    moonSize: 70,                           // 月亮尺寸
    sloganSizeNormal: 20,                   // 正常模式下的标语尺寸
    sloganSizeFlip: 16,                     // 翻转模式下的标语尺寸
    meteorStarWidth: 4,                     // 流星宽度
    meteorStarHeight: 100,                  // 流星高度
    meteorStarRotation: 30,                 // 流星基础倾斜角度
    meteorStarStartPosition: {              // 流星出现范围(设备长度百分比)
        min: 0.6,
        max: 0.9
    },
    meteorStarAnimateDuration: 2000,        // 流星动画执行时间
    decorateStarNum: 1 / 4000,              // 背景星星密度(a/b，每b平方像素内a个)
    decorateStarN: {                        // 背景星星边数范围
        min: 6,
        max: 8
    },
    decorateStarR: {                        // 背景星星外半径范围
        min: 2,
        max: 3.5
    },
    decorateStarR0: {                       // 背景星星内半径范围
        min: 1,
        max: 1.5
    },
    decorateStarFill: {                     // 背景星星填充色范围
        min: 185,
        max: 255
    },
    decorateStarOpacity: {                  // 背景星星透明度范围
        min: 0.4,
        max: 0.7
    },
    messageStarN: {                         // 寄语星星边数范围  
        min: 4,
        max: 4
    },
    messageStarR: {                         // 寄语星星外半径范围
        min: 5,
        max: 7
    },
    messageStarR0: {                        // 寄语星星内半径范围
        min: 0.5,
        max: 1.5
    },
    messageStarFill: {                      // 寄语星星填充色范围
        min: 144,
        max: 255
    },
    messageStarOpacity: {                   // 寄语星星透明度范围
        min: 0.7,
        max: 1
    },
    messageStarNum: 1 / 10000,              // 寄语星星密度(a/b，每b平方像素内a个)
    messageStarAnimateDuration: 5000,       // 寄语星星动画执行时间
}