

import React, { Component } from 'react'
import ReactDOM from 'react-dom';
import 'antd/dist/antd.css';
import '../../index.css';
import './Dashboard.less';
import InstanceList from './Components/InstanceList.js';
import { Layout, Menu, Breadcrumb, Button} from 'antd';

const { Header, Footer, Content } = Layout;


const mypageStyle={
    margin: "1rem",
    float: "right",
}

export default class Dashboard extends Component {
    render() {
        return (
            <div>
                <Layout className="layout">
                    <Header>
                    <div className="logo" />
                    <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
                        <Menu.Item key="1">nav 1</Menu.Item>
                        <Menu.Item key="2">nav 2</Menu.Item>
                        <Menu.Item key="3">nav 3</Menu.Item>
                        <Button danger style={mypageStyle}>비밀번호수정</Button>
                    </Menu>
                    </Header>
                    <Content style={{ padding: '0 50px' }}>
                    <InstanceList/>
                    </Content>
                    <Footer style={{ textAlign: 'center' }}>Ant Design ©2018 Created by Ant UED</Footer>
                </Layout>
            </div>
        )
    }
}