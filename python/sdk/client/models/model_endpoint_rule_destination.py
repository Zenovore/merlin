# coding: utf-8

"""
    Merlin

    API Guide for accessing Merlin's model management, deployment, and serving functionalities  # noqa: E501

    OpenAPI spec version: 0.14.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""

import pprint
import re  # noqa: F401

import six

class ModelEndpointRuleDestination(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """
    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'version_endpoint_id': 'str',
        'version_endpoint': 'VersionEndpoint',
        'weight': 'int'
    }

    attribute_map = {
        'version_endpoint_id': 'version_endpoint_id',
        'version_endpoint': 'version_endpoint',
        'weight': 'weight'
    }

    def __init__(self, version_endpoint_id=None, version_endpoint=None, weight=None):  # noqa: E501
        """ModelEndpointRuleDestination - a model defined in Swagger"""  # noqa: E501
        self._version_endpoint_id = None
        self._version_endpoint = None
        self._weight = None
        self.discriminator = None
        if version_endpoint_id is not None:
            self.version_endpoint_id = version_endpoint_id
        if version_endpoint is not None:
            self.version_endpoint = version_endpoint
        if weight is not None:
            self.weight = weight

    @property
    def version_endpoint_id(self):
        """Gets the version_endpoint_id of this ModelEndpointRuleDestination.  # noqa: E501


        :return: The version_endpoint_id of this ModelEndpointRuleDestination.  # noqa: E501
        :rtype: str
        """
        return self._version_endpoint_id

    @version_endpoint_id.setter
    def version_endpoint_id(self, version_endpoint_id):
        """Sets the version_endpoint_id of this ModelEndpointRuleDestination.


        :param version_endpoint_id: The version_endpoint_id of this ModelEndpointRuleDestination.  # noqa: E501
        :type: str
        """

        self._version_endpoint_id = version_endpoint_id

    @property
    def version_endpoint(self):
        """Gets the version_endpoint of this ModelEndpointRuleDestination.  # noqa: E501


        :return: The version_endpoint of this ModelEndpointRuleDestination.  # noqa: E501
        :rtype: VersionEndpoint
        """
        return self._version_endpoint

    @version_endpoint.setter
    def version_endpoint(self, version_endpoint):
        """Sets the version_endpoint of this ModelEndpointRuleDestination.


        :param version_endpoint: The version_endpoint of this ModelEndpointRuleDestination.  # noqa: E501
        :type: VersionEndpoint
        """

        self._version_endpoint = version_endpoint

    @property
    def weight(self):
        """Gets the weight of this ModelEndpointRuleDestination.  # noqa: E501


        :return: The weight of this ModelEndpointRuleDestination.  # noqa: E501
        :rtype: int
        """
        return self._weight

    @weight.setter
    def weight(self, weight):
        """Sets the weight of this ModelEndpointRuleDestination.


        :param weight: The weight of this ModelEndpointRuleDestination.  # noqa: E501
        :type: int
        """

        self._weight = weight

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(ModelEndpointRuleDestination, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, ModelEndpointRuleDestination):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
