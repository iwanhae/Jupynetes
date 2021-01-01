

import React, { Component } from 'react'
import 'antd/dist/antd.css';
import '../../index.css';
import { Button, Modal, Form, Input} from 'antd';
import CSS from 'csstype';

const mypageStyle: CSS.Properties ={
    margin: '1rem',
    float: 'right',
}

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

const onFinish = (values: String) => {
    console.log('Success:', values);
};

const onFinishFailed = () => {
};

const ChangePassword = () => {

    const [visible, setVisible] = React.useState(false);
    const [confirmLoading, setConfirmLoading] = React.useState(false);
    const [modalText, setModalText] = React.useState('Content of the modal');

    const showModal = () => {
      setVisible(true);
    };

    const handleOk = () => {
      setModalText('The modal will be closed after two seconds');
      setConfirmLoading(true);
      setTimeout(() => {
        setVisible(false);
        setConfirmLoading(false);
      }, 2000);
    };

    const handleCancel = () => {
      console.log('Clicked cancel button');
      setVisible(false);
    };

    return (
        <div>
                  <Button danger style={mypageStyle} onClick={showModal}>비밀번호수정</Button>
                  <Modal
                    title="비밀번호 수정"
                    visible={visible}
                    onOk={handleOk}
                    confirmLoading={confirmLoading}
                    onCancel={handleCancel}
                  >
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
                            label="아이디"
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
                                label="기존 비밀번호"
                                name="passwordOriginal"
                                rules={[
                                {
                                    required: true,
                                    message: 'Please input your password!',
                                },
                                ]}>
                                <Input.Password />
                            </Form.Item>
                            <Form.Item
                                label="변경할 비밀번호"
                                name="passwordNew"
                                rules={[
                                {
                                    required: true,
                                    message: 'Please input your password!',
                                },
                                ]}>
                                <Input.Password />
                            </Form.Item>
                            <Form.Item
                                label="비밀번호 재입력"
                                name="passwordRetype"
                                rules={[
                                {
                                    required: true,
                                    message: 'Please input your password!',
                                },
                                ]}>
                                <Input.Password />
                            </Form.Item>

                            <Form.Item {...tailLayout}>
                                <Button type="primary" htmlType="submit">
                                Submit
                                </Button>
                            </Form.Item>
                    </Form>


                  </Modal>
        </div>
    );
}

export default ChangePassword;