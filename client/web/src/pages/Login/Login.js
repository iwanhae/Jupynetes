import {React, Component} from 'react';
import ReactDOM from 'react-dom';
import 'antd/dist/antd.css';
import '../../index.css';
import './Login.less';

import { Typography, Layout, Space} from 'antd';
import LoginForm from "./Components/LoginForm.js";

const { Title } = Typography;
const { Content } = Layout;

export default class Login extends Component {

    render() {
        return (
            <div >
                <Layout >
                    <Space direction="vertical">
                        <Content >
                            <Title >Jupynetes</Title>
                            <div style={{height: '3rem'}}></div>
                            <LoginForm />
                        </Content>
                    </Space>
                </Layout>
            </div>
        )
    }
}