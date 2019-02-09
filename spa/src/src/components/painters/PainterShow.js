import React, { Component } from 'react'
import { PainterContext } from '../../AppContexts'
import { Link } from 'react-router-dom'


const PainterBox = ({Id, Name, CityOfOrigin})=>
    <div class="painter-box box" style= {{width:"600px"}}>
        <p>Fullname: {Name}</p>
        <p>City: {CityOfOrigin}</p>
        <Link class="button is-link" to={`/painter/${Id}/paints`} style= {{margin:"auto"}}> His/Her Paintings!</Link>
    </div>

class PainterShow extends Component {

    render(){
        return(
            <div className="painters-box">
            <PainterContext.Consumer>{
                painterContext => painterContext.paintersList.filter((painterJSON) => 
                    painterJSON.Id == this.props.match.params.painterId
                ).map((painterJSON)=>
                     <PainterBox {...painterJSON}/>                
                )
            }</PainterContext.Consumer>
            </div>             
        )
    }
}


export default PainterShow
