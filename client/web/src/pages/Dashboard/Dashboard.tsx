

import React, { Component } from 'react'
import 'antd/dist/antd.css';
import '../../index.css';
import './Dashboard.less';
import InstanceList from './Components/InstanceList.js';
import ChangePassword from '../ChangePassword/ChangePassword';
import { Layout, Menu, Button, Modal, Form, Input} from 'antd';
import CSS from 'csstype';

const { Header, Footer, Content } = Layout;
const mypageStyle: CSS.Properties ={
    margin: '1rem',
    float: 'right',
}

const Dashboard = () => {
    return (
        <>
            <Layout className="layout">
                <Header>
                <div className="logo" />
                <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
                    <Menu.Item key="1">nav 1</Menu.Item>
                    <Menu.Item key="2">nav 2</Menu.Item>
                    <Menu.Item key="3">nav 3</Menu.Item>
                    <Button danger style={mypageStyle}>로그아웃</Button>
                    <div style={mypageStyle}><ChangePassword></ChangePassword></div>
                </Menu>
                </Header>
                <Content style={{ padding: '0 50px' }}>
                <InstanceList/>
                </Content>
                <Footer style={{ textAlign: 'center' }}>Ant Design ©2018 Created by Ant UED</Footer>
            </Layout>
        </>
    );
}

export default Dashboard;