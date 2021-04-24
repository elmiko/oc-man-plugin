# oc man plugin

This is an experimental plugin for the [OpenShift command line tool](https://github.com/openshift/oc)
to provide a [man page](https://en.wikipedia.org/wiki/Man_page) like experience.

# How to

## Use

Type `oc man help` to see the detailed help page.

Type `oc man topics` to see a list of topics

Type `oc man <topic>`, replacing `<topic>` with the search term, to see the man
page for that topic.

## Build

**Prerequisites**
* Go lang 1.16+
* Python 3
* PyYaml

Type `make all` to gather the man pages and compile the `oc-man` binary. The output
will be in the `bin` directory.

A containerized build is available for environments where the prerequisites are not
met. Type `./hack/container-run.sh make` to build the binary.

Type `make clean` to remove the binary and downloaded content files.

## Install

As with [kubectl plugin extensions](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/),
place the binary `oc-man` file into the path of yourt host.

## Add a new man page

The [topics.yaml](topics.yaml) file contains the list of man page that will be downloaded
and compiled into the binary. They are organized as a list of title and location pairs,
where the title is a contiguous string and the location is a url. During the build process
the page titles will be matched with the content downloaded from the location stored
within the binary.

Add a new entry in the form

```yaml
- title: my-new-topic
  location: https://my-cdn.io/my-new-topic.md
```

To get started writing a man page, please see the [oc-man-page-template.md](oc-man-page-template.md)
file for further instructions.
