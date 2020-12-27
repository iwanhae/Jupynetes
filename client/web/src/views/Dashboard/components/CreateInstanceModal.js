import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import DialogTitle from '@material-ui/core/DialogTitle';
import Dialog from '@material-ui/core/Dialog';
import { blue } from '@material-ui/core/colors';
import GridItem from "components/Grid/GridItem.js";
import GridList from '@material-ui/core/GridList';
import GridListTile from '@material-ui/core/GridListTile';
import TextField from '@material-ui/core/TextField';

const useStyles = makeStyles({
  avatar: {
    backgroundColor: blue[100],
    color: blue[600],
  }, root: {
     width: "100px"
  }
});

CreateInstanceModal.propTypes = {
  onClose: PropTypes.func.isRequired,
  open: PropTypes.bool.isRequired,
  selectedValue: PropTypes.string.isRequired,
};

export default function CreateInstanceModal(props) {
  const classes = useStyles();
  const { onClose, selectedValue, open } = props;

  const handleClose = () => {
    onClose(selectedValue);
  };

  const handleListItemClick = (value) => {
    onClose(value);
  };

  return (
    <Dialog onClose={handleClose} aria-labelledby="서버 생성" open={open} maxWidth={'md'} fullWidth={true}>
      <DialogTitle id="simple-dialog-title">서버 생성</DialogTitle>
      <GridList>
        <GridListTile>
          <a>이름</a>
          <TextField
              label="name"
              id="instance-name"
           />
        </GridListTile>
        <GridListTile>
          <a>설명</a>
          <TextField
              label="name"
              id="instance-name"
           />
        </GridListTile>
        <GridListTile>
          <a>템플릿</a>
          <TextField
              label="name"
              id="instance-name"
           />
        </GridListTile>
        <GridListTile>
          <a>CPU</a>
          <TextField
              label="name"
              id="instance-name"
           />
        </GridListTile>
        <GridListTile>
          <a>Memory</a>
          <TextField
              label="name"
              id="instance-name"
           />
        </GridListTile>
                <GridListTile>
        <a>Nvidia_GPU</a>
          <TextField
              label="name"
              id="instance-name"
           />
        </GridListTile>
      </GridList>
    </Dialog>
  );
}