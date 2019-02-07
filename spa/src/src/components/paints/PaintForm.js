import React, { Component } from 'react'
import { PaintContext } from '../../AppContexts'
import { PainterContext } from '../../AppContexts'
import { Redirect } from 'react-router-dom'



class PaintForm extends Component{
    constructor(props) {
        super(props);
        this.state = {
          name: '',
          currentLocation: '',
          painter: 0
        };
    
        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleLocationChange = this.handleLocationChange.bind(this);
        this.handlePainterChange = this.handlePainterChange.bind(this);
    }

    handleNameChange(event) {
        this.setState({name: event.target.value});
    }

    handleLocationChange(event) {
        this.setState({currentLocation: event.target.value});
    }

    handlePainterChange(event) {
        alert(event.target.value)
        this.setState({painter: event.target.value});
    }
          

    render() {
        if(this.state.redirect){
            return <Redirect to="/" />
        }

        return (
    <PainterContext.Consumer>
    {painterContext =>        
    <PaintContext.Consumer>
    {paintContext =>
            <div class="field">
                <legend>Create New Painting Form</legend>
                <label class="label">Name</label>
                <div class="control">
                    <input class="input" type="text" placeholder="Name" value={this.state.name}  onChange={this.handleNameChange}/>
                </div>

                <label class="label">Current Location</label>
                <div class="control">
                    <input class="input" type="text" placeholder="Current Location" value={this.state.currentLocation} onChange={this.handleLocationChange}/>
                </div>        

                <div class="field">
                <label class="label">Painter</label>
                <div class="control">
                    <div class="select">
                    <select onChange={this.handlePainterChange}>
                        <option>Choose Painter</option>
                        {painterContext.paintersList.map((painter)=><option value={painter.Id}>{painter.Name}</option>)}
                    </select>
                    </div>
                </div>
                </div>                

                <div class="field is-grouped">
                    <div class="control">
                        <input class="button is-link" type="submit" onClick={paintContext.submitPaintForm.bind(this)} value="Submit" />
                    </div>
                    <div class="control">
                        <button class="button is-text">Cancel</button>
                    </div>
                </div>
            </div>
    }
    </PaintContext.Consumer>
    }
    </PainterContext.Consumer>
        )
    }
}

export default PaintForm