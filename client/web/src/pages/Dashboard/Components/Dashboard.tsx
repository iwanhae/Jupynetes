

import React, { Component } from 'react'
import 'antd/dist/antd.css';
import '../../../index.css';
import './Dashboard.less';
import ChangePassword from '../../ChangePassword/ChangePassword';
import { Layout, Button, Row, Col} from 'antd';
import CSS from 'csstype';
import { InstanceListContainer } from '../Container/InstanceListContainer';

const { Header, Footer, } = Layout;
const mypageStyle: CSS.Properties ={
    margin: '0 auto',
    float: 'right',
    marginRight: '5px'
}

const Dashboard = () => {
    return (
        <>
            <Layout className="layout">
                <Header>
                <div className="logo" />
                  <Row style={mypageStyle}>
                    <Col span={12}>
                      <ChangePassword></ChangePassword>
                    </Col>
                    <Col span={12}>
                      <Button danger >로그아웃</Button>
                    </Col>
                  </Row>
                </Header>
                <InstanceListContainer/>
                <Footer style={{ textAlign: 'center' }}>Ant Design ©2018 Created by Ant UED</Footer>
            </Layout>
        </>
    );
}

export default Dashboard;