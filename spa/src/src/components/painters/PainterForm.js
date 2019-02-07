import React, { Component } from 'react'
import { PainterContext } from '../../AppContexts'
import { Redirect } from 'react-router-dom'



class PainterForm extends Component{
    constructor(props) {
        super(props);
        this.state = {
          name: '',
          cityoforigin: '',
          redirect: false
        };
    
        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleCityChange = this.handleCityChange.bind(this);
    }

    handleNameChange(event) {
        this.setState({name: event.target.value});
    }

    handleCityChange(event) {
        this.setState({cityoforigin: event.target.value});
    }

          

    render() {
        if(this.state.redirect){
            return <Redirect to="/" />
        }

        return (
    <PainterContext.Consumer>
        {painterContext =>
            <div class="field">
                <legend>Create New Painter Form</legend>
                <label class="label">Name</label>
                <div class="control">
                    <input class="input" type="text" placeholder="Name" value={this.state.name}  onChange={this.handleNameChange}/>
                </div>

                <label class="label">City of Origin</label>
                <div class="control">
                    <input class="input" type="text" placeholder="City" value={this.state.cityoforigin} onChange={this.handleCityChange}/>
                </div>        

                <div class="field is-grouped">
                    <div class="control">
                        <input class="button is-link" type="submit" onClick={painterContext.submitPainterForm.bind(this)} value="Submit" />
                    </div>
                    <div class="control">
                        <button class="button is-text">Cancel</button>
                    </div>
                </div>
            </div>
        }
    </PainterContext.Consumer>
    
        )
    }
}

export default PainterForm