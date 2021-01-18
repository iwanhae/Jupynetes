import React, { FC } from "react";
import { ServerObject } from "../../../model/ServerObject";
import Instance from "../Components/Instance";


interface InstanceContainerProps {
  serverData: ServerObject
}

const InstanceContainer: FC<InstanceContainerProps> = ({serverData}) =>  {
    return <Instance serverData={serverData}></Instance>;
}

export default InstanceContainer;