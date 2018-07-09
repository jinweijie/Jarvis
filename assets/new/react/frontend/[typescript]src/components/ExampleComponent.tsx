import * as React from 'react';

import { IExampleContainerProps } from 'types';

export interface IExampleComponentStateProps {
  text: string;
}

export interface IExampleComponentDispatchProps {
  setText(): void;
}

type IExampleComponentProps = IExampleComponentStateProps & IExampleComponentDispatchProps & IExampleContainerProps;

export class ExampleComponent extends React.Component<IExampleComponentProps> {
  public onClick = (e: React.MouseEvent<HTMLAnchorElement>) => {
    e.preventDefault();
    this.props.setText();
  }

  public render() {
    return (
      <div className='center-align'>
        <div>
          <a href='#' className='btn todo-filter-btn waves-effect waves-light' onClick={this.onClick}>
            Append Text
              </a>
        </div>
        <div>
          {this.props.text}
        </div>
      </div>
    );
  }
}
