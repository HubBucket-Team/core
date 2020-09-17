<img width="240px" src="img/logo.png">

![build](https://img.shields.io/github/workflow/status/onepanelio/core/Publish%20dev%20docker%20image/master?color=01579b)
![code](https://img.shields.io/codacy/grade/d060fc4d1ac64b85b78f85c691ead86a?color=01579b)
[![release](https://img.shields.io/github/v/release/onepanelio/core?color=01579b)](https://github.com/onepanelio/core/releases)
[![sdk](https://img.shields.io/pypi/v/onepanel-sdk?color=01579b&label=sdk)](https://pypi.org/project/onepanel-sdk/)
[![docs](https://img.shields.io/github/v/release/onepanelio/core?color=01579b&label=docs)](https://docs.onepanel.io)
[![issues](https://img.shields.io/github/issues-raw/onepanelio/core?color=01579b&label=issues)](https://github.com/onepanelio/core/issues)
[![chat](https://img.shields.io/badge/support-slack-01579b)](https://onepanel-ce.slack.com/join/shared_invite/zt-eyjnwec0-nLaHhjif9Y~gA05KuX6AUg#/)
[![license](https://img.shields.io/github/license/onepanelio/core?color=01579b)](https://opensource.org/licenses/Apache-2.0)

Production scale, Kubernetes-native vision AI platform, with fully integrated components for model building, automated labeling, data processing and model training pipelines.

## Why Onepanel?

-  End-to-end workflow and infrastructure automation for production scale vision AI
-  Automatic resource management and on-demand scaling of CPU and GPU nodes
-  Easily scale your data processing and training pipelines to multiple nodes
-  Collaborate on all your deep learning tools and workflows through a unified web interface and SDKs
-  Scalability, flexibility and resiliency of Kubernetes without the deployment and configuration complexities

## Features
<table>
  <tr>
    <td width="50%" align="center">
      <h3>Image and video annotation with automatic annotation</h3>
      <img width="100%" src="img/auto-annotation.gif">
      <p>
        Annotate images and video with automatic annotation of bounding boxes and polygon masks, integrated with training pipelines to iteratively improve models for pre-annotation and inference
      </p>
    </td>
    <td width="50%" align="center">
      <h3>JupyterLab with TensorFlow, PyTorch and GPU support</h3>
      <img width="100%" src="img/jupyterlab.gif">
      <p>
        JupyterLab configured with extensions for debugging, Git/GitHub, notebook diffing and TensorBoard and support for Conda, OpenCV, Tensorflow and PyTorch with GPU and <a href="https://github.com/onepanelio/templates/tree/master/workspaces/jupyterlab">much more</a>
      </p>
    </td>
  </tr>
  <tr>
    <td width="50%" align="center">
      <h3>Auto scaling, distributed and parallel data processing and training pipelines</h3>
      <img width="100%" src="img/pipelines.gif">
      <p>
        Build fully reproducible, distributed and parallel data processing and training pipelines with real-time logs and output snapshots
      </p>
    </td>
    <td width="50%" align="center">
      <h3>Version controlled pipelines and environments as code</h3>
      <img width="100%" src="img/tools.gif">
      <p>
        Bring your own IDEs, annotation tools and pipelines with a version controlled YAML and Docker based template engine
      </p>
    </td>
  </tr>
</table>

-  Track and visualize model metrics and experiments with TensorBoard or bring your own experiment tracking tools.
-  Access and share tools like AirSim, Carla, Gazebo or OpenAI Gym through your browser with VNC enabled workspaces.
-  Extend Onepanel with powerful REST APIs and SDKs to further automate your pipelines and environments.
-  Workflows, environments and infrastructure are all defined as code and version controlled, making them reproducible and portable.
-  Powered by Kubernetes so you can deploy anywhere Kubernetes can run.

## Quick start
See [quick start guide](https://docs.onepanel.ai/docs/getting-started/quickstart) to get started with the platform of your choice.

### Quick start videos
[Getting started with Microsoft Azure](https://youtu.be/CQBIYfBk3Zk)\
[Getting started with Amazon EKS](https://youtu.be/Ipdd8f6D6IM)\
[Getting started with Google GKE](https://youtu.be/pZRO63SnQ8A)

## Community
See [documentation](https://docs.onepanel.ai) to get started or for more detailed operational and user guides.

To submit a feature request, report a bug or documentation issue, please open a GitHub [pull request](https://github.com/onepanelio/core/pulls) or [issue](https://github.com/onepanelio/core/issues).

For help, questions, release announcements and contribution discussions, join us on [Slack](https://join.slack.com/t/onepanel-ce/shared_invite/zt-eyjnwec0-nLaHhjif9Y~gA05KuX6AUg).

## Contributing

Onepanel is modular and consists of the following repositories:

[Core API](https://github.com/onepanelio/core/) (this repository) - Code base for backend (Go)\
[Core UI](https://github.com/onepanelio/core-ui/) - Code base for UI (Angular + TypeScript)\
[CLI](https://github.com/onepanelio/cli/) - Code base for Go CLI for installation and management (Go)\
[Manifests](https://github.com/onepanelio/core-ui/) - Kustomize manifests used by CLI for installation and management (YAML)\
[Python SDK](https://github.com/onepanelio/python-sdk/) - Python SDK code and documentation\
[Templates](https://github.com/onepanelio/templates) - Various Workspace, Workflow, Task and Sidecar Templates\
[Documentation](https://github.com/onepanelio/core-docs/) - The repository for documentation site\
[API Documentation](https://github.com/onepanelio/core-api-docs/) - API documentation if you choose to use the API directly

See `CONTRIBUTING.md` in each repository for development guidelines. Also, see [contribution guide](https://docs.onepanel.ai/docs/getting-started/contributing) for additional guidelines.


## Acknowledgments
Onepanel seamlessly integrates the following excellent open source projects. We are grateful for the support these communities provide and do our best to contribute back as much as possible.

[Argo](https://github.com/argoproj/argo)\
[CVAT](https://github.com/opencv/cvat)\
[JupyterLab](https://github.com/jupyterlab/jupyterlab)


## License
Onepanel is licensed under [Apache 2.0](https://github.com/onepanelio/core/blob/master/LICENSE).

## Need a managed solution?
Visit our [website](https://www.onepanel.io/) for more information about our managed offerings.
