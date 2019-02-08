import React, { Component } from 'react'
import { PaintContext } from '../../AppContexts'
import { Link } from 'react-router-dom'

const PaintBox = ({Name, CurrentLocation, Image})=>
    <div class="">
        <div>
            <h3>Paiting Name: {Name}</h3>
            <span>Current Location: {CurrentLocation}</span>
        </div>
        <img src={`data:image/png;base64, ${Image}`} alt="Red dot" />
    </div>


class PaintsList extends Component {

    render(){
        return(
            <div>
                <PaintContext.Consumer>{
                    paintContext => paintContext.paintsList.filter((paintJSON)=>
                        paintJSON.PainterId == this.props.match.params.painterId
                    ).map((paintJSON) =>   
                           <PaintBox {...paintJSON} />
                    )
                }</PaintContext.Consumer>
                <Link to="/paint/create" class="button is-primary">Add new Painting</Link> 
            </div>
        )
    }
}

export default PaintsList