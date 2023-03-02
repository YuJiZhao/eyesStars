# eyesStars

网站地址：http://stars.eyescode.top

代码在对应分支里。

# 一：介绍
![在这里插入图片描述](https://img-blog.csdnimg.cn/2df93cd93b6b4521a1913f79ee293037.png)

我开发这个网站的契机是一次机器学习的作业，当时我在查论文，偶然间进入了爱墙网，这个网站是用来表达爱意的，可惜这个网站现在已经关停了。当时我便被这新奇的想法所吸引，网站的关停更是激励我开发一个类似的。

本人母胎单身至今，所以自然不想只用来表达爱意，毕竟这跟我没什么关系。所以我开发了耶瞳星空，星空里闪烁的星星，都寄托着一句话，这句话可以是任何类型的(违fa除外)。总之，当我们仰望星空时想说什么，星星里就可以是什么。

通过点击闪烁的四角星星，可以看到其他人说的话，点击月亮，可以看到网站的介绍，这也是功能入口，可以从月亮进入发布弹窗，发布属于自己的星星。

# 二：技术相关
我开发这个网站还有一个重要的原因，就是为了试验我的用户中心。我目前除了耶瞳星空，还有一个个人网站：耶瞳空间（http://space.eyescode.top）。我不想让两个网站的用户数据隔离开，因此我需要一个服务管理我所有应用的用户数据，然后我开发了用户中心。现在用户中心负责网站的认证和授权，以及用户基础数据的管理。所以现在只要在用户中心登录一次，再使用我的网站的话只要跳转授权就行。

技术栈：
+ 前端：Nuxt3 + ZRender
+ 后端：Gin + Thrift + Nacos
+ 数据库：MySQL

这个网站很简单，所以技术细节就没必要说了，大家可以去github看看代码，注释也有很多。
