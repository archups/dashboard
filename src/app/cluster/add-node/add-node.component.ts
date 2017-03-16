import {Component, OnInit, Input} from "@angular/core";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {CustomValidators} from "ng2-validation";
import {ApiService} from "../../api/api.service";
import {ClusterModel} from "../../api/model/ClusterModel";
import {CreateNodeModel} from "../../api/model/CreateNodeModel";
import {NodeInstanceFlavors} from "../../api/model/NodeProviderConstants";

@Component({
  selector: 'kubermatic-add-node',
  templateUrl: './add-node.component.html',
  styleUrls: ['./add-node.component.scss']
})

export class AddNodeComponent implements OnInit {
  @Input() cluster: any;
  @Input() clusterName: string;
  @Input() seedDcName: string;
  @Input() nodeProvider: string;

  public addNodeForm: FormGroup;
  public clusterModel: ClusterModel;
  public createNodeModel: CreateNodeModel;
  public nodeDcName: string;
  public node: any;
  public nodeSpec: any = {spec: {}}
  public nodeInstances: number = 1;
  public nodeSizes: any;
  public sshKeys: any;


  constructor(private api: ApiService, private formBuilder: FormBuilder) { }

  ngOnInit() {

    this.nodeDcName = this.cluster.spec.cloud.dc;
    this.nodeProvider = this.cluster.dc.spec.provider;
    this.nodeSpec.spec.dc = this.cluster.spec.cloud.dc;

    this.getProviderNodeSpecification();

    this.addNodeForm = this.formBuilder.group({
      node_count: [1, [<any>Validators.required, CustomValidators.min(1), CustomValidators.max(20)]],
      node_size: ['', [<any>Validators.required]]
    });
  }

  public getProviderNodeSpecification() {
    switch (this.nodeProvider) {
      case 'aws' : {
        this.nodeSizes = NodeInstanceFlavors.AWS;
        return this.nodeSizes;
      }

      case 'digitalocean' : {
        this.api.getDigitaloceanSizes(this.cluster.spec.cloud.digitalocean.token).subscribe(result => {
            this.nodeSizes = result.sizes;
            return this.nodeSizes;
          }
        );
      }
    }
  }

  public setProviderNodeSpecification(): void {
    this.nodeInstances = this.addNodeForm.controls["node_count"].value;
    switch (this.nodeProvider) {
      case 'aws' : {
        this.nodeSpec.spec.aws = {
          type: this.addNodeForm.controls["node_size"].value
        };
        return;
      }

      case 'digitalocean' : {
        this.nodeSpec.spec.digitalocean = {
          sshKeys: this.cluster.spec.cloud.digitalocean.sshKeys,
          size: this.addNodeForm.controls["node_size"].value
        };
        return;
      }

      default : {
        break;
      }
    }
  }

  public addNode(): void {
    this.setProviderNodeSpecification();
    this.clusterModel = new ClusterModel(this.seedDcName, this.clusterName);
    this.createNodeModel = new CreateNodeModel(this.nodeInstances,this.nodeSpec.spec);

    this.api.createClusterNode(this.clusterModel, this.createNodeModel).subscribe(result => {
      this.node = result;
    })
  }
}
