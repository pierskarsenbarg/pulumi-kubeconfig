# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


VERSION = "0.0.0"
def readme():
    try:
        with open('README.md', encoding='utf-8') as f:
            return f.read()
    except FileNotFoundError:
        return "kubeconfig Pulumi Package - Development Version"


setup(name='pierskarsenbarg_pulumi_kubeconfig',
      python_requires='>=3.8',
      version=VERSION,
      description="Kubeconfig provider",
      long_description=readme(),
      long_description_content_type='text/markdown',
      packages=find_packages(),
      package_data={
          'pierskarsenbarg_pulumi_kubeconfig': [
              'py.typed',
              'pulumi-plugin.json',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'pulumi>=3.0.0,<4.0.0',
          'pulumi_azure_native>=2.0.0,<3.0.0',
          'pulumi_kubernetes>=4.0.0,<5.0.0',
          'pulumni_aws>=6.0.0,<7.0.0',
          'semver>=2.8.1'
      ],
      zip_safe=False)
