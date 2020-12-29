import React from 'react';
import { Typography, Layout, Form, Input, Button, Checkbox} from 'antd';

import 'antd/dist/antd.css';
import '../../../index.css';
import './LoginForm.less';

const { Title } = Typography;
const layout = {
  labelCol: {
    span: 6,
  },
  wrapperCol: {
    span: 12,
  },
};
const tailLayout = {
  wrapperCol: {
    offset: 6,
    span: 12,
  },
};

const onFinish = (values) => {
    console.log('Success:', values);
};

const onFinishFailed = (errorInfo) => {
    console.log('Failed:', errorInfo);
};

export default function LoginForm(props) {
    return (
        <div style={{background: 'rgba(255, 255, 0, 1)', width: '30%', margin: "0 auto"}}>
            <div style={{backgorund: 'rgba(255, 0, 0, 1)', width: "80%", margin: "0 auto"}}>
                <Title>Login</Title>
                <Form
                    {...layout}
                    name="basic"
                    initialValues={{
                        remember: true,
                    }}
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}
                >
                            <Form.Item
                            label="Username"
                            name="username"
                            rules={[
                            {
                                required: true,
                                message: 'Please input your username!',
                            },
                            ]}>
                            <Input />
                            </Form.Item>
                            <Form.Item
                                label="Password"
                                name="password"
                                rules={[
                                {
                                    required: true,
                                    message: 'Please input your password!',
                                },
                                ]}>
                                <Input.Password />
                            </Form.Item>
                            <Form.Item {...tailLayout} name="remember" valuePropName="checked">
                                <Checkbox>Remember me</Checkbox>
                            </Form.Item>

                            <Form.Item {...tailLayout}>
                                <Button type="primary" htmlType="submit">
                                Submit
                                </Button>
                            </Form.Item>
                </Form>
            </div>
        </div>
    );
}
