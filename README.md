# mood
个人博客  
**技术栈**：gin + gorm + mysql + redis + vue  

**功能**：简洁的个人博客。一个比较大的模块文章管理，功能包括：增删改查、权限管理、限流。文本编辑使用markdown，使用vue组件。
目的就是简洁、高效的记录文字，没有考虑很多花里胡哨的功能。权限部分，该项目属于个人分享内容的地方，不存在其他用户，所以使用一个全局的密钥验证即可。
初衷是记录自己的笔记和一些随想，所以分类就直接硬编码为技术和随想，也没有标签一说，没必要搞太多的东西。文字编辑打算在本地环境进行，然后通过上传md文件的形式。
将文件解析且存储起来，更新和添加就上传文件解析即可。

**接口**：文章列表、文章详情、文章权限、文章置顶、文章更新、添加文章、获取权限、文章归档(只使用get和post)  
[get] blog/list  
[get] blog/detail/:id     
[get] blog/archive  
[post] blog/detail/:id   
[post] blog/top  
[post] blog/add  
[post] blog/authority  
[post] blog/upload

**数据库**：blog、secret  
id creat_time update_time delete_time  
title  #blog标题
top  #是否置顶   
image  #blog封面图  
content  #内容  
md_path  #md文件地址  
music_url  #音乐地址  
category  #分类 1、技术  2、随想  
key  #密钥内容  
role  #密钥权限  1、可读私密文章 2、管理员权限  

以上就是全部功能，以后可能会拓展功能，如评论、笔记本（如将leetcode刷题笔记、经典文档翻译）。
