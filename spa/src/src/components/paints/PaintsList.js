import React, { Component } from 'react'
import { PaintContext } from '../../AppContexts'
import { Link } from 'react-router-dom'
import './PaintsList.css'

const PaintBox = ({Name, CurrentLocation, Image})=>
    <div class="paint-box">
        <div>
            <h3>Paiting Name: {Name}</h3>
            <span>Current Location: {CurrentLocation}</span>
        </div>
        <Link to="/">
            <img src={`data:image/png;base64, ${Image}`} style={{maxHeight: "140px"}} />
        </Link>
    </div>


class PaintsList extends Component {

    render(){
        return(
            <div>
                <div class="paints-list-box">
                    <PaintContext.Consumer>{
                        paintContext => paintContext.paintsList.filter((paintJSON)=>
                            paintJSON.PainterId == this.props.match.params.painterId
                        ).map((paintJSON) =>   
                            <PaintBox {...paintJSON} />
                        )
                    }</PaintContext.Consumer>
                </div>
            <Link to="/paint/create" class="button is-primary">Add new Painting</Link> 
            </div>
        )
    }
}

export default PaintsList